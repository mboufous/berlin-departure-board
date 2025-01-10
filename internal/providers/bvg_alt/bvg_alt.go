package bvg_alt

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/mboufous/berlin-departure-board/internal"
)

const (
	ProviderName = "BVG-ALT"
	baseUrl      = "https://v6.bvg.transport.rest"
)

type BVGALTProvider struct {
	internal.BaseProvider
}

func NewProvider(client *http.Client) *BVGALTProvider {
	return &BVGALTProvider{
		internal.BaseProvider{
			BaseUrl: baseUrl,
			Client:  client,
		},
	}
}

// HealthCheck performs a health check on the BVG-ALT API
func (p *BVGALTProvider) HealthCheck(ctx context.Context) (internal.APIProviderStatus, error) {
	healthUrl := fmt.Sprintf("%s/health", p.BaseUrl)
	slog.Debug("Performing health check", "url", healthUrl)

	var response HealthCheckResponse
	requestOptions := internal.RequestOptions{
		Method: http.MethodGet,
		Target: &response,
	}
	if err := p.Do(ctx, requestOptions); err != nil {
		return internal.APIProviderStatus{}, fmt.Errorf("failed to perform health check: %w", err)
	}

	slog.Info("Health check completed", "url", healthUrl, "up", response.Ok)
	return internal.APIProviderStatus{Up: response.Ok}, nil
}

// GetStation fetches a station by its ID
func (p *BVGALTProvider) GetStation(ctx context.Context, stationID string) (internal.Station, error) {
	stationUrl, err := NewURLBuilder(p.BaseUrl, stationID).
		WithParam("linesOfStops", "true").
		WithParam("language", "en").
		Build()
	if err != nil {
		return internal.Station{}, fmt.Errorf("failed to build station URL: %w", err)
	}
	slog.Debug("Fetching station", "url", stationUrl, "stationID", stationID)

	var response StationResponse
	requestOptions := internal.RequestOptions{
		Method: http.MethodGet,
		Target: &response,
	}
	if err := p.Do(ctx, requestOptions); err != nil {
		slog.Error("Failed to fetch station", "url", stationUrl, "stationID", stationID, "error", err)
		return internal.Station{}, fmt.Errorf("failed to fetch station: %w", err)
	}

	slog.Info("Successfully fetched station", "url", stationUrl, "stationID", stationID)

	return internal.Station{
		ID:    response.Id,
		Name:  response.Name,
		Lines: mapLines(response.Lines),
	}, nil
}

// GetDepartures fetches the departures for a station by its ID
func (p *BVGALTProvider) GetDepartures(ctx context.Context, stationID string, params internal.APIDeparturesParams) (internal.DepartureBoard, error) {
	builder := NewURLBuilder(p.BaseUrl, stationID).
		WithEndpoint("/departures").
		WithDefaultParam("duration", params.Duration, "30").
		WithDefaultParam("results", params.MaxResultCount, "4").
		WithDefaultParam("linesOfStops", params.LinesOfStops, "false").
		WithDefaultParam("remarks", params.Remarks, "true").
		WithParam("language", "en").
		WithParam("direction", params.Direction).
		WithProductFilters(params.Products)

	departuresUrl, err := builder.Build()
	if err != nil {
		return internal.DepartureBoard{}, fmt.Errorf("failed to build departures URL: %w", err)
	}

	slog.Debug("Fetching departures", "url", departuresUrl, "stationID", stationID)

	var response DepartureResponse
	requestOptions := internal.RequestOptions{
		Method: http.MethodGet,
		Target: &response,
	}
	if err := p.Do(ctx, requestOptions); err != nil {
		slog.Error("Failed to fetch departures", "url", departuresUrl, "stationID", stationID, "error", err)
		return internal.DepartureBoard{}, fmt.Errorf("failed to fetch departures: %w", err)
	}

	slog.Info("Successfully fetched departures", "url", departuresUrl, "stationID", stationID)

	return internal.DepartureBoard{
		Departures:  mapDepartures(response),
		LastUpdated: time.Now(),
	}, nil
}

func mapLines(lines []LinesResponse) []internal.Product {
	products := make([]internal.Product, len(lines))
	for i, l := range lines {
		products[i] = internal.Product{
			Name: l.Name,
			Type: l.ProductName,
		}
	}
	return products
}

func (p *BVGALTProvider) Name() string {
	return ProviderName
}

func mapDepartures(response DepartureResponse) []internal.Departure {
	departures := make([]internal.Departure, len(response.Departures))
	for i, d := range response.Departures {
		departures[i] = internal.Departure{
			ID:          d.TripId,
			When:        d.When,
			PlannedWhen: d.PlannedWhen,
			Delay:       d.Delay,
			Product: internal.Product{
				Name: d.Line.Name,
				Type: d.Line.Product,
			},
			Direction: d.Direction,
			Destination: internal.Station{
				ID:   d.Destination.Id,
				Name: d.Destination.Name,
			},
			Hints:     mapRemarks(d.Remarks),
			Platform:  d.Platform,
			Occupancy: d.Occupancy,
		}
	}
	return departures
}

func mapRemarks(response []RemarkResponse) []internal.Remark {
	remarks := make([]internal.Remark, len(response))
	for i, r := range response {
		remarks[i] = internal.Remark{
			Type:   r.Type,
			Header: r.Summary,
			Body:   r.Text,
		}
	}
	return remarks
}
