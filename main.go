package main

import (
	"context"
	"fmt"
	"github.com/mboufous/berlin-departure-board/hafas"
	"github.com/mboufous/berlin-departure-board/transportproviders/bvg"
	"log"
	"log/slog"
	"os"
	"time"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	}))
	slog.SetDefault(logger)

	client := hafas.NewClient(&bvg.Provider{})

	ctx := context.WithValue(context.Background(), "User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")

	station, err := client.Station.Get(ctx, hafas.StationParams{
		StationID: "900008102",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Station:", station.Name)

	departureBoard, err := client.Departure.Get(ctx, hafas.DepartureParams{
		Station:         station,
		When:            time.Now(),
		ProductsFilter:  hafas.NewProductFilter().AddProduct(hafas.ProductSubway).Build(),
		ShowRemarks:     true,
		DurationMinutes: 20,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(departureBoard)
}
