package bvg_alt

import "time"

type ProductsResponse struct {
	Suburban bool `json:"suburban"`
	Subway   bool `json:"subway"`
	Tram     bool `json:"tram"`
	Bus      bool `json:"bus"`
	Ferry    bool `json:"ferry"`
	Express  bool `json:"express"`
	Regional bool `json:"regional"`
}

type StationResponse struct {
	Type     string           `json:"type"`
	Id       string           `json:"id"`
	Name     string           `json:"name"`
	Products ProductsResponse `json:"products"`
	Lines    []LinesResponse  `json:"lines"`
}

type LinesResponse struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Public      bool   `json:"public"`
	ProductName string `json:"productName"`
	Mode        string `json:"mode"`
	Product     string `json:"product"`
}
type DepartureResponse struct {
	Departures []struct {
		TripId          string          `json:"tripId"`
		Stop            StationResponse `json:"stop"`
		When            time.Time       `json:"when"`
		PlannedWhen     *time.Time      `json:"plannedWhen"`
		Delay           int             `json:"delay"`
		Platform        *string         `json:"platform"`
		PlannedPlatform *string         `json:"plannedPlatform"`
		PrognosisType   string          `json:"prognosisType"`
		Direction       string          `json:"direction"`
		Provenance      interface{}     `json:"provenance"`
		Line            struct {
			Type        string `json:"type"`
			Id          string `json:"id"`
			FahrtNr     string `json:"fahrtNr"`
			Name        string `json:"name"`
			Public      bool   `json:"public"`
			AdminCode   string `json:"adminCode"`
			ProductName string `json:"productName"`
			Mode        string `json:"mode"`
			Product     string `json:"product"`
			Operator    struct {
				Type string `json:"type"`
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"operator"`
		} `json:"line"`
		Remarks     []RemarkResponse `json:"remarks"`
		Origin      interface{}      `json:"origin"`
		Destination struct {
			Type     string `json:"type"`
			Id       string `json:"id"`
			Name     string `json:"name"`
			Location struct {
				Type      string  `json:"type"`
				Id        string  `json:"id"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"location"`
			Products ProductsResponse `json:"products"`
		} `json:"destination"`
		CurrentTripPosition struct {
			Type      string  `json:"type"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"currentTripPosition"`
		Occupancy string `json:"occupancy"`
	} `json:"departures"`
	RealtimeDataUpdatedAt int `json:"realtimeDataUpdatedAt"`
}

type RemarkResponse struct {
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Text    string `json:"text"`
}

type HealthCheckResponse struct {
	Ok bool `json:"ok"`
}
