package bvg_alt

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBVGALTProvider_GetDepartures(t *testing.T) {
	// Define the mock JSON response as a variable
	mockResponse := `{
  "departures": [
    {
      "tripId": "1|97794|0|86|10122024",
      "stop": {
        "type": "stop",
        "id": "900100001",
        "name": "S+U Friedrichstr. Bhf (Berlin)",
        "location": {
          "type": "location",
          "id": "900100001",
          "latitude": 52.520519,
          "longitude": 13.386448
        },
        "products": {
          "suburban": true,
          "subway": true,
          "tram": true,
          "bus": true,
          "ferry": false,
          "express": false,
          "regional": true
        }
      },
      "when": "2024-12-10T22:28:00+01:00",
      "plannedWhen": "2024-12-10T21:39:00+01:00",
      "delay": 2940,
      "platform": "1",
      "plannedPlatform": "1",
      "prognosisType": "prognosed",
      "direction": "Eisenhüttenstadt, Bahnhof",
      "provenance": null,
      "line": {
        "type": "line",
        "id": "odre1-re1",
        "fahrtNr": "73828",
        "name": "RE1",
        "public": true,
        "adminCode": "ODRE1_",
        "productName": "RE",
        "mode": "train",
        "product": "regional",
        "operator": {
          "type": "operator",
          "id": "odeg-ostdeutsche-eisenbahn-gmbh",
          "name": "ODEG Ostdeutsche Eisenbahn GmbH"
        }
      },
      "remarks": [
        {
          "type": "hint",
          "code": "FK",
          "text": "Bicycle conveyance"
        },
        {
          "type": "hint",
          "code": "ib",
          "text": "Fahrradmitnahme leicht gemacht: www.vbb.de/radimregio"
        },
        {
          "id": "252045",
          "type": "warning",
          "summary": "Information.",
          "text": "Die Fahrplanauskunft im DB Navigator sowie auf www.bahn.de ist bis zum 10.12.2024 einschließlich nicht aktuell. Grund hierfür in ein technischer Fehler. Bitte nutzen Sie die Fahrplanauskunft des vbb unter <a href=\"https://www.vbb.de/fahrinfo/\" target=\"_blank\">https://www.vbb.de/fahrinfo/</a> oder in der vbb App.",
          "icon": {
            "type": "HIM0",
            "title": null
          },
          "priority": 32,
          "products": {
            "suburban": false,
            "subway": false,
            "tram": false,
            "bus": false,
            "ferry": false,
            "express": false,
            "regional": true
          },
          "company": "VBB",
          "categories": [
            0
          ],
          "validFrom": "2024-12-06T00:00:00+01:00",
          "validUntil": "2024-12-10T23:59:00+01:00",
          "modified": "2024-12-05T15:07:04+01:00"
        }
      ],
      "origin": null,
      "destination": {
        "type": "stop",
        "id": "900311307",
        "name": "Eisenhüttenstadt, Bahnhof",
        "location": {
          "type": "location",
          "id": "900311307",
          "latitude": 52.1478,
          "longitude": 14.658323
        },
        "products": {
          "suburban": false,
          "subway": true,
          "tram": false,
          "bus": true,
          "ferry": false,
          "express": false,
          "regional": true
        }
      }
    }
  ],
  "realtimeDataUpdatedAt": 1733863990
}`

	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Create the provider using the mock server URL
	provider := NewProvider(&http.Client{})
	provider.BaseUrl = mockServer.URL

	// Call GetDepartures
	ctx := context.Background()
	departureBoard, err := provider.GetDepartures(ctx, "900100001")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, departureBoard)
	assert.Equal(t, 1, len(departureBoard.Departures))
	assert.Equal(t, "1|97794|0|86|10122024", departureBoard.Departures[0].ID)
	assert.Equal(t, "RE1", departureBoard.Departures[0].Product.Name)
	assert.Equal(t, "Eisenhüttenstadt, Bahnhof", departureBoard.Departures[0].Destination.Name)
	assert.Equal(t, []string{"subway", "bus", "regional"}, departureBoard.Departures[0].Destination.Products)
}

func TestBVGALTProvider_GetStation(t *testing.T) {
	mockResponse := `{
		"type": "station",
		"id": "900100001",
		"name": "S+U Friedrichstr. Bhf (Berlin)",
		"products": {
			"suburban": true,
			"subway": true,
			"tram": true,
			"bus": true,
			"ferry": false,
			"express": false,
			"regional": true
		}
	}`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/stops/900100001", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	provider := NewProvider(&http.Client{})
	provider.BaseUrl = mockServer.URL

	ctx := context.Background()
	station, err := provider.GetStation(ctx, "900100001")

	assert.NoError(t, err)
	assert.NotNil(t, station)
	assert.Equal(t, "900100001", station.ID)
	assert.Equal(t, "S+U Friedrichstr. Bhf (Berlin)", station.Name)
	assert.Equal(t, []string{"suburban", "subway", "tram", "bus", "regional"}, station.Products)
}
