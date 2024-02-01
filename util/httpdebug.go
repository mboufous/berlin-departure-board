package util

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
)

type LoggingTransport struct{}

func (s *LoggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	reqBytes, _ := httputil.DumpRequestOut(r, true)

	resp, err := http.DefaultTransport.RoundTrip(r)

	respBytes, _ := httputil.DumpResponse(resp, true)

	slog.Debug("HTTP", slog.String("Request", string(reqBytes)), slog.String("Response", string(respBytes)))

	return resp, err
}
