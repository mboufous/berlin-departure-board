package bvg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDepartureResponse_ValidResponse(t *testing.T) {
	bvgProvider := &Provider{}

	body := loadMockedBVGResponse(t, "valid_departure_response.json")
	defer body.Close()

	departureBoard, err := bvgProvider.ParseDepartureResponse(body, true)

	assert.NoError(t, err, "Expected no error for valid JSON response")
	assert.NotEmpty(t, departureBoard, "Departure board should not be empty for valid JSON")
	assert.NotEmpty(t, departureBoard.Departures, "Departures should not be empty for valid JSON")
	assert.NotEmpty(t, departureBoard.Remarks, "Remarks should not be empty for valid JSON when enabled")

	assert.Len(t, departureBoard.Departures, 7, "Unexpected number of departures")
	assert.Len(t, departureBoard.Remarks, 1, "Unexpected number of remarks")

	testDeparture := departureBoard.Departures[0]
	assert.Equal(t, "U Reinickendorfer Str. (Berlin)", testDeparture.Stop.Name)
	assert.Equal(t, "900008102", testDeparture.Stop.ID)
	assert.Equal(t, "Märkisches Viertel, Wilhelmsruher Damm", testDeparture.Direction)
	assert.Equal(t, "120", testDeparture.Line.Name)
	assert.Equal(t, "210500", testDeparture.When.Format(timeLayout))
	assert.Equal(t, 4, testDeparture.Delay)

	testDepartureWithDayOffset := departureBoard.Departures[1]
	assert.Equal(t, "010300", testDepartureWithDayOffset.When.Format(timeLayout))
	assert.Equal(t, "20240124", testDepartureWithDayOffset.When.Format(dateLayout))

	testRemark := departureBoard.Remarks[0]
	assert.Equal(t, "Rail strike until Monday, 29 January, 6 p.m.", testRemark.Header)

}

func TestParseDepartureResponse_InvalidResponse(t *testing.T) {
	bvgProvider := &Provider{}

	body := loadMockedBVGResponse(t, "invalid_departure_response.json")
	defer body.Close()

	departureBoard, err := bvgProvider.ParseDepartureResponse(body, true)

	assert.Error(t, err, "Expected an error for invalid JSON response")
	assert.Empty(t, departureBoard, "Departure board should be empty for invalid JSON")

}
