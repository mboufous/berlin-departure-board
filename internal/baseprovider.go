package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type BaseProvider struct {
	BaseUrl string
	Client  *http.Client
}

const (
	ProductSuburban = "suburban"
	ProductSubway   = "subway"
	ProductTram     = "tram"
	ProductBus      = "bus"
	ProductFerry    = "ferry"
	ProductExpress  = "express"
	ProductRegional = "regional"
)

type APIDeparturesParams struct {
	// Duration       string //Show departures for how many minutes
	MaxResultCount int  //Max number of results to return
	LinesOfStops   bool //lines of each stop/station
	Remarks        bool
	Direction      string
	Products       []string
}

type RequestOptions struct {
	Method  string
	Headers map[string]string
	Body    any
	Target  any
}

func (p *BaseProvider) Do(ctx context.Context, opts RequestOptions) error {
	var bodyReader io.Reader
	if opts.Body != nil {
		bodyBytes, err := json.Marshal(opts.Body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}
	req, err := http.NewRequestWithContext(ctx, opts.Method, p.BaseUrl, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	resp, err := p.Client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) || errors.Is(err, http.ErrHandlerTimeout) {
			return fmt.Errorf("request timed out: %w", err)
		}
		return fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	slog.Debug("HTTP Response", "url", p.BaseUrl, "status", resp.StatusCode)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	if opts.Target != nil {
		if err := json.NewDecoder(resp.Body).Decode(opts.Target); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}
	return nil
}
