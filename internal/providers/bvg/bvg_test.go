package bvg_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/mboufous/berlin-departure-board/internal"
	"github.com/mboufous/berlin-departure-board/internal/providers/bvg"
)

// MockTransport implements http.RoundTripper for testing
type MockTransport struct {
	Response *http.Response
	Err      error
}

func (m *MockTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return m.Response, m.Err
}

// loadMockResponse is a helper function to load test data
func loadMockResponse(filePath string) (*http.Response, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(data)),
		Header:     make(http.Header),
	}, nil
}

// setupTest creates a test provider with mock response
func setupTest(t *testing.T, responseFile string) *bvg.Provider {
	t.Helper()
	mockResponse, err := loadMockResponse(responseFile)
	if err != nil {
		t.Fatalf("Failed to load mock response: %v", err)
	}

	client := &http.Client{
		Transport: &MockTransport{Response: mockResponse},
	}

	return bvg.NewProvider(client)
}

func Test_BVG_GetDepartures(t *testing.T) {
	t.Run("default departure board response", func(t *testing.T) {
		provider := setupTest(t, "testdata/departures_api_response_default.json")
		board, err := provider.GetDepartures(context.Background(), "900100001", internal.APIDeparturesParams{})

		if err != nil {
			t.Fatalf("GetDepartures() error = %v", err)
		}

		// Common validations
		if time.Since(board.LastUpdated) > time.Minute {
			t.Errorf("Expected LastUpdated to be recent, got '%v'", board.LastUpdated)
		}
		if board.Station.ID != "900100001" {
			t.Errorf("Expected station ID to be '900100001', got '%s'", board.Station.ID)
		}

		if len(board.Departures) == 0 {
			t.Fatal("Expected departures to be non-empty")
		}

		firstDeparture := board.Departures[0]

		// Basic departure info
		if firstDeparture.ID != "1|301|0|86|9012025" {
			t.Errorf("Expected ID to be '1|301|0|86|9012025', got '%s'", firstDeparture.ID)
		}
		if firstDeparture.Direction != "S+U Gesundbrunnen Bhf (Berlin)" {
			t.Errorf("Expected direction to be 'S+U Gesundbrunnen Bhf (Berlin)', got '%s'", firstDeparture.Direction)
		}

		// Product details
		if firstDeparture.Product.Type != "S" {
			t.Errorf("Expected product type to be 'S', got '%s'", firstDeparture.Product.Type)
		}

		// Platform and status
		if firstDeparture.Platform != "12" {
			t.Errorf("Expected platform to be '12', got '%s'", firstDeparture.Platform)
		}
		if firstDeparture.Canceled {
			t.Error("Expected departure to not be canceled")
		}

		// Additional info
		if len(firstDeparture.Hints) != 0 {
			t.Errorf("Expected no remarks, got %d remarks", len(firstDeparture.Hints))
		}
		if firstDeparture.Occupancy != "low occupancy expected" {
			t.Errorf("Expected occupancy to be 'low occupancy expected', got '%s'", firstDeparture.Occupancy)
		}

		if firstDeparture.Product.Name != "S1" {
			t.Errorf("Expected product name to be 'S1', got '%s'", firstDeparture.Product.Name)
		}
	})

	t.Run("canceled departure board response", func(t *testing.T) {
		provider := setupTest(t, "testdata/departures_api_response_canceled.json")
		board, err := provider.GetDepartures(context.Background(), "900100001", internal.APIDeparturesParams{})

		if err != nil {
			t.Fatalf("GetDepartures() error = %v", err)
		}

		// Common validations
		if time.Since(board.LastUpdated) > time.Minute {
			t.Errorf("Expected LastUpdated to be recent, got '%v'", board.LastUpdated)
		}
		if board.Station.ID != "900100001" {
			t.Errorf("Expected station ID to be '900100001', got '%s'", board.Station.ID)
		}

		if len(board.Departures) == 0 {
			t.Fatal("Expected departures to be non-empty")
		}

		firstDeparture := board.Departures[0]

		// Basic departure info
		if firstDeparture.ID != "1|78824|0|86|9012025" {
			t.Errorf("Expected ID to be '1|78824|0|86|9012025', got '%s'", firstDeparture.ID)
		}
		if firstDeparture.Direction != "Flughafen BER" {
			t.Errorf("Expected direction to be 'Flughafen BER', got '%s'", firstDeparture.Direction)
		}

		// Platform and status
		if firstDeparture.Platform != "1" {
			t.Errorf("Expected platform to be '1', got '%s'", firstDeparture.Platform)
		}
		if !firstDeparture.Canceled {
			t.Error("Expected departure to be canceled")
		}

		// Remarks for canceled train
		if len(firstDeparture.Hints) == 0 {
			t.Error("Expected remarks for canceled train")
		}

		if firstDeparture.Occupancy != "low occupancy expected" {
			t.Errorf("Expected occupancy to be 'low occupancy expected', got '%s'", firstDeparture.Occupancy)
		}

		if firstDeparture.Product.Type != "RE" {
			t.Errorf("Expected product type to be 'RE', got '%s'", firstDeparture.Product.Type)
		}

		if firstDeparture.Product.Name != "RE8" {
			t.Errorf("Expected product name to be 'RE8', got '%s'", firstDeparture.Product.Name)
		}

		if firstDeparture.Warnings[0].Header != "Bauvorankündigung" {
			t.Errorf("Expected warning to be 'Bauvorankündigung', got '%s'", firstDeparture.Warnings[0].Header)
		}
	})
}

func Test_BVG_GetStation(t *testing.T) {

	provider := setupTest(t, "testdata/station_api_response.json")
	station, err := provider.GetStation(context.Background(), "900008102")
	if err != nil {
		t.Fatalf("GetStation() error = %v", err)
	}

	if station.ID != "900008102" {
		t.Errorf("Expected station ID to be '900008102', got '%s'", station.ID)
	}

	if station.Name != "U Reinickendorfer Str. (Berlin)" {
		t.Errorf("Expected station name to be 'U Reinickendorfer Str. (Berlin)', got '%s'", station.Name)
	}

	if len(station.Lines) != 6 {
		t.Errorf("Expected 6 lines, got %d", len(station.Lines))
	}

	if station.Lines[0].Type != "U" {
		t.Errorf("Expected first line type to be 'U', got '%s'", station.Lines[0].Type)
	}
	if station.Lines[0].Name != "U6" {
		t.Errorf("Expected first line name to be 'U6', got '%s'", station.Lines[0].Name)
	}
}
