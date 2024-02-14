package bvg

import (
	"github.com/mboufous/berlin-departure-board/hafas"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDepartureResponse_ValidResponse(t *testing.T) {
	bvgProvider := &APIProvider{}

	body, _ := loadMockedBVGResponse("valid_departure_response.json")
	defer body.Close()

	departureBoard, err := bvgProvider.ParseDepartureResponse(body)
	assert.NoError(t, err, "Expected no error for valid JSON response")
	assert.NotEmpty(t, departureBoard.Lines, "Expected a full list of lines")
	assert.Len(t, departureBoard.Lines, 3)

	var firstLine hafas.Line
	for _, line := range departureBoard.Lines {
		if line.Product.Name == "U6" {
			firstLine = line
		}
	}

	assert.Len(t, firstLine.Directions, 2)

	var firstDirection hafas.Direction
	for _, direction := range firstLine.Directions {
		if direction.Name == "Alt-Mariendorf" {
			firstDirection = direction
		}
	}
	assert.Len(t, firstDirection.Departures, 12)
}

// BenchmarkConvertLines-10    	  172573	      6850 ns/op	    8704 B/op	     185 allocs/op
func BenchmarkConvertLines(b *testing.B) {
	bvgProvider := &APIProvider{}

	body, _ := loadMockedBVGResponse("valid_departure_response.json")
	defer body.Close()
	apiResult, _ := bvgProvider.ParseBaseResponse(body)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bvgProvider.convertLines(apiResult.Res)
	}

	b.ReportAllocs()
}

func TestParseDepartureResponse_InvalidResponse(t *testing.T) {
	bvgProvider := &APIProvider{}

	body, _ := loadMockedBVGResponse("invalid_departure_response.json")
	defer body.Close()

	departureBoard, err := bvgProvider.ParseDepartureResponse(body)

	assert.Error(t, err, "Expected an error for invalid JSON response")
	assert.Empty(t, departureBoard, "Departure board should be empty for invalid JSON")

}
