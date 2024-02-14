package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mboufous/berlin-departure-board/cache"
	"github.com/mboufous/berlin-departure-board/cache/adapters"
	"github.com/mboufous/berlin-departure-board/cache/encoders"
	"github.com/mboufous/berlin-departure-board/hafas"
	"github.com/mboufous/berlin-departure-board/transportproviders/bvg"
	"log"
	"log/slog"
	"os"
	"time"
)

const (
	defaultExpirationTime  = 2 * time.Minute
	defaultCleanupInterval = 10 * time.Minute
)

// TODO: another a 2nd transport provider and round robing between the 2 apis
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
	}))
	slog.SetDefault(logger)

	store := adapters.NewMemoryAdapter[[]byte](defaultExpirationTime, defaultCleanupInterval)
	appCache := cache.NewCache(store, encoders.NewGobEncoder())

	client := hafas.NewClient(&bvg.APIProvider{}, hafas.WithCache(appCache))

	ctx := context.WithValue(context.Background(), "User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")

	// Get station details + get available Products
	station, err := client.Station.Get(ctx, hafas.StationParams{
		StationID: "900008102",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Station: %s\n", station.Name)
	for _, p := range station.Products {
		fmt.Println(p.Name)
	}

	departureBoard, err := client.Departure.Get(ctx, hafas.DepartureParams{
		Station:          station.ID,
		ProductsFilter:   hafas.NewLineFilter("U", "B"),
		DirectionsFilter: hafas.NewLineFilter("S+U Pankow"),
	})

	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(departureBoard, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
