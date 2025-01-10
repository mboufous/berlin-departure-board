package bvg

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/mboufous/berlin-departure-board/internal"
)

const (
	ProviderName = "BVG"
	baseUrl      = "https://bvg-apps-ext.hafas.de/bin/mgate.exe"
	dateLayout   = "20060102"
	timeLayout   = "150405"
)

type Provider struct {
	internal.BaseProvider
}

// NewProvider initializes a new BVGProvider with the provided HTTP client.
func NewProvider(client *http.Client) *Provider {
	return &Provider{
		BaseProvider: internal.BaseProvider{
			BaseUrl: baseUrl,
			Client:  client,
		},
	}
}

// GetDepartures retrieves departure information for a specific station.
func (p *Provider) GetDepartures(ctx context.Context, stationID string, params internal.APIDeparturesParams) (internal.DepartureBoard, error) {
	requestBody := createDepartureBoardRequest(stationID, params)

	var response BVGApiResponse
	requestOptions := internal.RequestOptions{
		Method: http.MethodPost,
		Body:   requestBody,
		Target: &response,
	}

	if err := p.Do(ctx, requestOptions); err != nil {
		return internal.DepartureBoard{}, fmt.Errorf("failed to fetch departures: %w", err)
	}

	if err := validateAPIResponse(response); err != nil {
		return internal.DepartureBoard{}, err
	}

	apiResult := response.SvcResL[0]
	if len(apiResult.Res.JnyL) == 0 {
		return internal.DepartureBoard{}, errors.New("no departures found in the response")
	}

	station := apiResult.Res.Common.LocL[0]

	return internal.DepartureBoard{
		Station: internal.Station{
			ID:    station.ExtId,
			Name:  station.Name,
			Lines: extractLines(apiResult.Res.Common.ProdL),
		},
		Departures:  extractDepartures(apiResult.Res),
		LastUpdated: time.Now(),
	}, nil

}

type JourneyMessages struct {
	Occupancy string
	Hints     []internal.Remark
	Warnings  []internal.Remark
}

func extractMessages(apiResult *SvcResData, journey *Journey) JourneyMessages {
	messages := JourneyMessages{}

	// stop-specific messages (Occupancy)
	for _, msg := range append(journey.StbStop.MsgL, journey.MsgL...) {
		if msg.RemX != nil {
			rem := apiResult.Common.RemL[*msg.RemX]
			if strings.HasPrefix(rem.Code, "text.occup.loc") {
				messages.Occupancy = rem.TxtN
			}
		}
	}

	// Process journey messages (Hints and Warnings)
	for _, msg := range journey.MsgL {
		if msg.RemX != nil {
			rem := apiResult.Common.RemL[*msg.RemX]
			if rem.Code != "FK" && !strings.HasPrefix(rem.Code, "text.occup.loc") {
				messages.Hints = append(messages.Hints, internal.Remark{
					Type:   msg.Type,
					Header: rem.Code,
					Body:   rem.TxtN,
				})
			}
		}
		if msg.HimX != nil {
			warning := apiResult.Common.HimL[*msg.HimX]
			messages.Warnings = append(messages.Warnings, internal.Remark{
				Type:   msg.Type,
				Header: warning.Head,
				Body:   warning.Text,
			})
		}
	}

	return messages
}

func extractDepartures(apiResult *SvcResData) []internal.Departure {
	var extractedDepartures []internal.Departure
	for _, journey := range apiResult.JnyL {
		departureTime, delay := extractDepartureTime(&journey)
		messages := extractMessages(apiResult, &journey)

		extractedDepartures = append(extractedDepartures, internal.Departure{
			ID:    journey.Jid,
			When:  departureTime,
			Delay: delay,
			Product: internal.Product{
				Name: apiResult.Common.ProdL[journey.StbStop.DProdX].ProdCtx.Line,
				Type: strings.TrimSpace(apiResult.Common.ProdL[journey.StbStop.DProdX].ProdCtx.CatOut),
			},
			Direction: journey.DirTxt,
			Destination: internal.Station{
				ID:   apiResult.Common.LocL[journey.ProdL[0].TLocX].ExtId,
				Name: apiResult.Common.LocL[journey.ProdL[0].TLocX].Name,
			},
			Platform:  journey.StbStop.DPlfS.Txt,
			Canceled:  journey.IsCncl,
			Occupancy: messages.Occupancy,
			Hints:     messages.Hints,
			Warnings:  messages.Warnings,
		})
	}
	return extractedDepartures
}

func extractDepartureTime(journey *Journey) (time.Time, int) {
	plannedDepartureTimeRaw := journey.StbStop.DTimeS
	newDepartureTimeRaw := journey.StbStop.DTimeR
	departureDate := journey.Date

	// Parse the planned departure time
	departureTime := parseDepartureTime(plannedDepartureTimeRaw, departureDate)
	delay := 0

	if isDepartureDelayed(journey) {
		newDepartureTime := parseDepartureTime(newDepartureTimeRaw, departureDate)
		delay = int(math.Round(newDepartureTime.Sub(departureTime).Minutes()))
		departureTime = newDepartureTime
	}

	return departureTime, delay
}

func isDepartureDelayed(journey *Journey) bool {
	return journey.StbStop.DTimeR != "" && journey.StbStop.DTimeS != journey.StbStop.DTimeR
}

func parseDepartureTime(timeStr, dateStr string) time.Time {
	// Load the Germany time zone
	location, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		slog.Warn("failed to load Germany time zone", "error", err)
		return time.Time{}
	}

	var offsetDays time.Duration
	adjustedTime := timeStr

	if isDayOffsetPresent(timeStr) {
		adjustedTime = timeStr[2:]
		offsetDays, err = time.ParseDuration(fmt.Sprintf("%sh", timeStr[0:2]))
		if err != nil {
			slog.Warn("failed to parse time offset", "error", err)
			return time.Time{}
		}
	}
	baseDate, err := time.ParseInLocation(dateLayout+timeLayout, dateStr+adjustedTime, location)
	if err != nil {
		slog.Warn("failed to parse date", "error", err)
		return time.Time{}
	}

	return baseDate.Add(offsetDays * 24)
}

func isDayOffsetPresent(timeStr string) bool {
	return len(timeStr) > len(timeLayout)
}

// GetStation fetches detailed information for a specific station by ID.
func (p *Provider) GetStation(ctx context.Context, stationID string) (internal.Station, error) {
	requestBody := createStationRequest(stationID)

	var response BVGApiResponse
	requestOptions := internal.RequestOptions{
		Method: http.MethodPost,
		Body:   requestBody,
		Target: &response,
	}

	if err := p.Do(ctx, requestOptions); err != nil {
		return internal.Station{}, fmt.Errorf("failed to fetch station: %w", err)
	}

	if err := validateAPIResponse(response); err != nil {
		return internal.Station{}, err
	}

	apiResult := response.SvcResL[0]
	if len(apiResult.Res.LocL) == 0 {
		return internal.Station{}, errors.New("no station found in the response")
	}

	station := apiResult.Res.LocL[0]
	return internal.Station{
		ID:    station.ExtId,
		Name:  station.Name,
		Lines: extractLines(apiResult.Res.Common.ProdL),
	}, nil
}

// HealthCheck performs a health check on the BVG API.
func (p *Provider) HealthCheck(ctx context.Context) (internal.APIProviderStatus, error) {
	requestBody := createServerStatusRequest()

	var response BVGApiResponse
	requestOptions := internal.RequestOptions{
		Method: http.MethodPost,
		Body:   requestBody,
		Target: &response,
	}

	if err := p.Do(ctx, requestOptions); err != nil {
		return internal.APIProviderStatus{}, fmt.Errorf("health check failed: %w", err)
	}

	if err := validateAPIResponse(response); err != nil {
		return internal.APIProviderStatus{}, err
	}
	return internal.APIProviderStatus{Up: true}, nil
}

// Name returns the name of the provider.
func (p *Provider) Name() string {
	return ProviderName
}

// String returns the string representation of the provider.
func (p *Provider) String() string {
	return p.Name()
}

// Helper Functions
func validateAPIResponse(response BVGApiResponse) error {
	if response.Err != "OK" {
		return fmt.Errorf("API error: %s", response.ErrTxt)
	}
	if len(response.SvcResL) == 0 {
		return errors.New("empty API response")
	}
	if response.SvcResL[0].Err != "OK" {
		return fmt.Errorf("API service error: %s", response.SvcResL[0].ErrTxt)
	}
	return nil
}

func extractLines(products []ProdL) []internal.Product {
	var lines []internal.Product
	for _, product := range products {
		lines = append(lines, internal.Product{
			Name: product.ProdCtx.Line,
			Type: strings.TrimSpace(product.ProdCtx.CatOut),
		})
	}
	return lines
}
