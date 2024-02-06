package hafas

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultDurationMinutes = 10
)

type DepartureService Service

type DepartureBoard struct {
	Departures []Departure
	Remarks    []Remark
}

type Departure struct {
	Stop      Station
	Direction string
	When      time.Time
	Delay     int
	Line      Line
}

type Remark struct {
	Header string
	Body   string
}

type Line struct {
	Name string
}

type DepartureParams struct {
	Station         *Station
	When            time.Time
	ProductsFilter  uint8
	DurationMinutes int
	ShowRemarks     bool
}

type DepartureProvider interface {
	NewDepartureRequest(param *DepartureParams) (*http.Request, error)
	ParseDepartureResponse(body io.ReadCloser, showRemarks bool) (*DepartureBoard, error)
}

func (s *DepartureService) Get(ctx context.Context, params any) (*DepartureBoard, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	departureParams, err := validateDepartureRequestParams(params)
	if err != nil {
		return nil, fmt.Errorf("params validation failed: %w", err)
	}

	req, err := s.client.provider.NewDepartureRequest(departureParams)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to get departureBoard: %w", err)
	}

	departureBoard, err := s.client.provider.ParseDepartureResponse(resp.Body, departureParams.ShowRemarks)
	if err != nil {
		return nil, fmt.Errorf("failed to parse departureBoard: %w", err)
	}

	return departureBoard, nil
}

func validateDepartureRequestParams(params any) (*DepartureParams, error) {

	if departureParams, ok := params.(DepartureParams); ok {
		if departureParams.Station == nil {
			return nil, errors.New("validation error: invalide station")
		}
		if departureParams.When.Before(time.Now().Add(-time.Hour)) {
			return nil, errors.New("validation error: cannot use a date in the past")
		}
		if departureParams.ProductsFilter <= 0 || departureParams.ProductsFilter > MaxProductsFilterBitmask {
			departureParams.ProductsFilter = MaxProductsFilterBitmask
		}
		if departureParams.DurationMinutes <= 0 {
			departureParams.DurationMinutes = DefaultDurationMinutes
		}
		return &departureParams, nil
	}
	return nil, errors.New("type assertion: params is not a departure request param type")
}
