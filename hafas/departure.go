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

type DepartureService Service

type DepartureBoard struct {
	Lines   []Line
	Remarks []Remark
}

type Line struct {
	Product    Product
	Directions []Direction
}

type Direction struct {
	Name       string
	Departures []Departure
	Remarks    []Remark
}

type Departure struct {
	When time.Time
}

type Remark struct {
	Header string
	Body   string
}

type DepartureParams struct {
	Station                      string
	MaxDeparturesDurationMinutes int
	DirectionsFilter             LineFilter
	ProductsFilter               LineFilter
}

type DepartureProvider interface {
	NewDepartureRequest(param *DepartureParams) (*http.Request, error)
	ParseDepartureResponse(body io.ReadCloser) (*DepartureBoard, error)
}

func (s *DepartureService) Get(ctx context.Context, params any) (*DepartureBoard, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	departureParams, err := validateDepartureRequestParams(params)
	if err != nil {
		return nil, fmt.Errorf("params validation failed: %w", err)
	}

	return s.getDepartureBoard(ctx, departureParams)

}

func (s *DepartureService) getDepartureBoard(ctx context.Context, params *DepartureParams) (*DepartureBoard, error) {
	if s.client.cache != nil {
		if departureBoard, err := s.getDepartureBoardFromCache(ctx, params); err == nil {
			slog.Debug("Cache hit for departure board", "cache_key", s.getCacheKey(params))
			departureBoard.Lines = s.filterLines(departureBoard.Lines, params)
			return departureBoard, nil
		} else if errors.Is(err, cache.ErrObjectNotFound) == false {
			slog.Error("Error accessing cache", "error", err)
		}
		slog.Debug("Cache miss for departure board", "cache_key", s.getCacheKey(params))
	}
	// We will always get all the departures for all the products to cache them, then filter the cached result.
	departureBoard, err := s.getDepartureBoardFromSource(ctx, params)
	if err != nil {
		return nil, err
	}

	if s.client.cache != nil {
		if err := s.cacheDepartureBoard(ctx, params, departureBoard); err != nil {
			slog.Warn("Failed to cache departure board", "station_id", params.Station, "error", err)
		}
	}
	departureBoard.Lines = s.filterLines(departureBoard.Lines, params) //TODO: are we changing the cached version or any other instance for other users
	return departureBoard, nil
}

func (s *DepartureService) getCacheKey(params *DepartureParams) string {
	return fmt.Sprintf("dep:%s", params.Station)
}

func (s *DepartureService) getDepartureBoardFromCache(ctx context.Context, params *DepartureParams) (*DepartureBoard, error) {
	var d DepartureBoard
	if err := s.client.cache.Get(ctx, s.getCacheKey(params), &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (s *DepartureService) cacheDepartureBoard(ctx context.Context, params *DepartureParams, board *DepartureBoard) error {
	return s.client.cache.Put(ctx, s.getCacheKey(params), board, s.getAdaptiveTTL(time.Now(), board.Lines))
}

func (s *DepartureService) getDepartureBoardFromSource(ctx context.Context, params *DepartureParams) (*DepartureBoard, error) {
	req, err := s.client.provider.NewDepartureRequest(params)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to get departureBoard: %w", err)
	}

	departureBoard, err := s.client.provider.ParseDepartureResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse departureBoard: %w", err)
	}

	// Cap to MaxDeparturesPerDirection
	//if len(departureBoard.Departures) > MaxDeparturesPerDirection {
	//	departureBoard.Departures = departureBoard.Departures[:MaxDeparturesPerDirection]
	//}
	return departureBoard, nil
}

func (s *DepartureService) filterLines(lines []Line, params *DepartureParams) []Line {
	var filteredLines []Line

	for _, line := range lines {
		if !params.ProductsFilter.Filter(line.Product.Type) {
			continue
		}

		var filteredDirections []Direction
		for _, direction := range line.Directions {
			if params.DirectionsFilter.Filter(direction.Name) {
				filteredDirections = append(filteredDirections, direction)
			}
		}

		// Only add the line if there are any directions that match the filter
		if len(filteredDirections) > 0 {
			filteredLine := Line{
				Product:    line.Product,
				Directions: filteredDirections,
			}
			filteredLines = append(filteredLines, filteredLine)
		}
	}

	return filteredLines
}

func validateDepartureRequestParams(params any) (*DepartureParams, error) {

	if departureParams, ok := params.(DepartureParams); ok {
		if _, err := strconv.Atoi(departureParams.Station); err != nil {
			return nil, fmt.Errorf("departure params validation failed: %w", err)
		}
		if departureParams.MaxDeparturesDurationMinutes < 0 {
			departureParams.MaxDeparturesDurationMinutes = 0
		}

		return &departureParams, nil
	}
	return nil, errors.New("type assertion: params is not a departure request param type")
}
