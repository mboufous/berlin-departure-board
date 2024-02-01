package main

import (
	"context"
	"fmt"
	"github.com/mboufous/berlin-departure-board/hafas"
	"github.com/mboufous/berlin-departure-board/transportproviders/bvg"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}))
	slog.SetDefault(logger)

	bvgProvider := bvg.NewProvider()
	httpClient := http.DefaultClient

	client := hafas.NewClient(bvgProvider, hafas.WithEnableDebugMode(), hafas.WithHTTPClient(httpClient))

	station, err := client.Station.Get(context.Background(), bvg.StationRequestParams{
		StationID: "900008102",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(station.Name)

	// TODO: Options
	// 	Max number of departures
	//	show hitns and warnings
	//
	departures, err := client.Departures.Get(context.Background(), station)
}
