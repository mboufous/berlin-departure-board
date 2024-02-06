package hafas

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type StationService Service

type Station struct {
	ID       string
	Name     string
	Products []Product
}

type StationParams struct {
	StationID string
}

type StationProvider interface {
	NewStationRequest(params *StationParams) (*http.Request, error)
	ParseStationResponse(body io.ReadCloser) (*Station, error)
}

func (s *StationService) Get(ctx context.Context, params any) (*Station, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	stationParams, err := validateStationRequestParams(params)
	if err != nil {
		return nil, fmt.Errorf("params validation failed: %w", err)
	}

	req, err := s.client.provider.NewStationRequest(stationParams)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to get station: %w", err)
	}

	station, err := s.client.provider.ParseStationResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing station response: %w", err)
	}
	return station, nil
}

func validateStationRequestParams(params any) (*StationParams, error) {
	if stationParams, ok := params.(StationParams); ok {
		if _, err := strconv.Atoi(stationParams.StationID); err != nil {
			return nil, fmt.Errorf("station id validation failed: %w", err)
		}
		return &stationParams, nil
	}
	return nil, errors.New("type assertion: params is not a station request param type")
}
