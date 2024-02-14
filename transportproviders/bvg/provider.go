package bvg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	apiURL         = "https://bvg-apps-ext.hafas.de/bin/mgate.exe"
	dateLayout     = "20060102"
	timeLayout     = "150405"
	fullDateLayout = dateLayout + timeLayout
)

type APIProvider struct {
}

func (p *APIProvider) newRequest(payload any) (*http.Request, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("payload marshal failed: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewReader(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("http request creation failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (p *APIProvider) ParseBaseResponse(responseBody io.ReadCloser) (*SvcRes, error) {
	bvgResponse := &ApiResponse{}
	err := json.NewDecoder(responseBody).Decode(bvgResponse)
	if err != nil {
		return nil, fmt.Errorf("bvg response decoding: %w", err)
	}

	if bvgResponse.Err != "OK" {
		return nil, fmt.Errorf("bvg api error: %s: %s", bvgResponse.Err, bvgResponse.ErrTxt)
	}

	if len(bvgResponse.SvcResL) == 0 {
		return nil, errors.New("empty data returned")
	}

	apiResult := &bvgResponse.SvcResL[0]

	if apiResult.Err != "OK" {
		if apiResult.Err == "LOCATION" {
			return nil, errors.New("invalid station id")
		}
		return nil, fmt.Errorf("API error: %s", apiResult.ErrTxt)
	}

	return apiResult, nil
}
