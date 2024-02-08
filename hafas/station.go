package hafas

import (
	"context"
	"errors"
	"fmt"
	"github.com/mboufous/berlin-departure-board/cache"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"
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

	return s.getStation(ctx, stationParams)
}

func (s *StationService) getStation(ctx context.Context, params *StationParams) (*Station, error) {
	if s.client.cache != nil {
		if station, err := s.getStationFromCache(ctx, params.StationID); err == nil {
			slog.Debug("Cache hit for station", "station_id", params.StationID, "station_name", station.Name)
			return station, nil
		} else if !errors.Is(err, cache.ErrObjectNotFound) {
			slog.Error("Error accessing cache", "error", err)
		}
		slog.Debug("Cache miss for station", "station_id", params.StationID)
	}

	station, err := s.getStationFromSource(ctx, params)
	if err != nil {
		return nil, err
	}

	if s.client.cache != nil {
		if err := s.cacheStation(ctx, params.StationID, station); err != nil {
			slog.Warn("Failed to cache station", "station_id", params.StationID, "error", err)
		}
	}

	return station, nil
}

func (s *StationService) getStationFromCache(ctx context.Context, stationID string) (*Station, error) {
	var station Station
	if err := s.client.cache.Get(ctx, s.getCacheKey(stationID), &station); err != nil {
		return nil, err
	}
	return &station, nil
}

func (s *StationService) cacheStation(ctx context.Context, stationID string, station *Station) error {
	return s.client.cache.Put(ctx, s.getCacheKey(stationID), station, 1*time.Hour)
}

func (s *StationService) getCacheKey(id string) string {
	return "station:" + id
}

func (s *StationService) getStationFromSource(ctx context.Context, params *StationParams) (*Station, error) {
	req, err := s.client.provider.NewStationRequest(params)
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
