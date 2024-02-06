package bvg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStationResponse_ValidResponse(t *testing.T) {
	bvgProvider := &Provider{}

	body := loadMockedBVGResponse(t, "valid_station_response.json")
	defer body.Close()

	station, err := bvgProvider.ParseStationResponse(body)

	assert.NoError(t, err, "Expected no error for valid JSON response")
	assert.NotEmpty(t, station, "Station should not be empty for valid JSON")

	assert.Equal(t, "U Reinickendorfer Str. (Berlin)", station.Name)
	assert.Equal(t, "900008102", station.ID)
	assert.Len(t, station.Products, 6)

}

func TestParseStationResponse_InvalidResponse(t *testing.T) {
	bvgProvider := &Provider{}

	body := loadMockedBVGResponse(t, "invalid_station_response.json")
	defer body.Close()

	station, err := bvgProvider.ParseStationResponse(body)

	assert.Error(t, err, "Expected an error for invalid JSON response")
	assert.Empty(t, station, "Departure board should be empty for invalid JSON")

}
