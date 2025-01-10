package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mboufous/berlin-departure-board/internal"
)

type Handler struct {
	client *internal.APIClient
}

func NewHandler(client *internal.APIClient) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthcheck", h.handleHealth)
	mux.HandleFunc("GET /station/{id}", h.handleGetStation)
	mux.HandleFunc("GET /station/{id}/departures", h.handleGetDepartures)

	// Add version prefix
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", mux))

	return v1
}

func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	status, err := h.client.HealthCheck(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func (h *Handler) handleGetStation(w http.ResponseWriter, r *http.Request) {
	stationId := r.PathValue("id")

	station, err := h.client.GetStation(r.Context(), stationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(station)
}

func (h *Handler) handleGetDepartures(w http.ResponseWriter, r *http.Request) {
	stationId := r.PathValue("id")
	params := internal.APIDeparturesParams{
		Products:  []string{r.URL.Query().Get("product")},
		Direction: r.URL.Query().Get("direction"),
		MaxResultCount: func() int {
			count, err := strconv.Atoi(r.URL.Query().Get("results"))
			if err != nil {
				return 4
			}
			return count
		}(),
		LinesOfStops: r.URL.Query().Get("linesOfStops") == "true",
		Remarks:      r.URL.Query().Get("remarks") == "true",
	}

	departures, err := h.client.GetDepartures(r.Context(), stationId, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(departures)
}
