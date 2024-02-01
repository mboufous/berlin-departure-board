package bvg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mboufous/berlin-departure-board/hafas"
	"io"
	"net/http"
	"strconv"
	"time"
)

type StationRequestParams struct {
	StationID string
}

type DepartureRequestParams struct {
	When      time.Time
	MaxResult int
}

type BVGProvider struct {
	apiURL string
}

func NewProvider() *BVGProvider {
	return &BVGProvider{
		apiURL: "https://bvg-apps-ext.hafas.de/bin/mgate.exe",
	}
}

func (p *BVGProvider) NewStationRequest(params any) (*http.Request, error) {
	stationParams, err := validateStationRequestParams(params)
	if err != nil {
		return nil, fmt.Errorf("params validation failed: %w", err)
	}
	payload := CreateBVGStationRequestPayload(stationParams.StationID)
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("payload marshal failed: %w", err)
	}

	r, err := http.NewRequest(http.MethodPost, p.apiURL, bytes.NewReader(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("http request creation failed: %w", err)
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 8_2_4) Gecko/20100101 Firefox/54.5")

	return r, nil
}

func (p *BVGProvider) NewStationResponse(body io.ReadCloser) (*hafas.Station, error) {
	stationResponse := new(BVGStationResponse)
	err := json.NewDecoder(body).Decode(stationResponse)
	if err != nil {
		return nil, err
	}

	if len(stationResponse.SvcResL) == 0 {
		return nil, errors.New("unknown error status")
	}

	apiResult := stationResponse.SvcResL[0]

	if apiResult.Err != "OK" {
		if apiResult.Err == "LOCATION" {
			return nil, errors.New("station not found")
		}
		return nil, fmt.Errorf("API error: %s", apiResult.ErrTxt)
	}

	if len(apiResult.Res.LocL) == 0 {
		return nil, errors.New("empty station details")
	}

	if apiResult.Meth != "LocDetails" {
		return nil, errors.New("empty station details")
	}

	// convert to app station data
	stationData := hafas.Station{
		ID:       apiResult.Res.LocL[0].ExtId,
		Name:     apiResult.Res.LocL[0].Name,
		Products: populateProducts(apiResult.Res.Common.ProdL),
	}

	return &stationData, nil
}

func (p *BVGProvider) NewDepartureRequest(params any) (*http.Request, error) {
	stationParams, err := validateStationRequestParams(params)
	if err != nil {
		return nil, fmt.Errorf("params validation failed: %w", err)
	}
	payload := CreateBVGStationRequestPayload(stationParams.StationID)
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("payload marshal failed: %w", err)
	}

	r, err := http.NewRequest(http.MethodPost, p.apiURL, bytes.NewReader(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("http request creation failed: %w", err)
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 8_2_4) Gecko/20100101 Firefox/54.5")

	return r, nil
}

func (p *BVGProvider) NewDepartureResponse(body io.ReadCloser) (*hafas.Departure, error) {
	stationResponse := new(BVGStationResponse)
	err := json.NewDecoder(body).Decode(stationResponse)
	if err != nil {
		return nil, err
	}

	if len(stationResponse.SvcResL) == 0 {
		return nil, errors.New("unknown error status")
	}

	apiResult := stationResponse.SvcResL[0]

	if apiResult.Err != "OK" {
		if apiResult.Err == "LOCATION" {
			return nil, errors.New("station not found")
		}
		return nil, fmt.Errorf("API error: %s", apiResult.ErrTxt)
	}

	if len(apiResult.Res.LocL) == 0 {
		return nil, errors.New("empty station details")
	}

	if apiResult.Meth != "LocDetails" {
		return nil, errors.New("empty station details")
	}

	// convert to app station data
	stationData := hafas.Station{
		ID:       apiResult.Res.LocL[0].ExtId,
		Name:     apiResult.Res.LocL[0].Name,
		Products: populateProducts(apiResult.Res.Common.ProdL),
	}

	return &stationData, nil
}

func validateStationRequestParams(params any) (*StationRequestParams, error) {
	if stationParams, ok := params.(StationRequestParams); ok {
		if _, err := strconv.Atoi(stationParams.StationID); err != nil {
			return nil, fmt.Errorf("station id validation failed: %w", err)
		}
		return &stationParams, nil
	}
	return nil, errors.New("type assertion: wrong params type")
}

func populateProducts(products []Product) []hafas.Product {
	var hafasProducts []hafas.Product

	for _, product := range products {
		hafasProducts = append(hafasProducts, hafas.Product{
			Name: product.NameS,
			Type: product.ProdCtx.CatOutS,
		})
	}

	return hafasProducts
}
