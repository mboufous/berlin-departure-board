package hafas

import (
	"context"
	"errors"
	"fmt"
	"github.com/mboufous/berlin-departure-board/util"
	"net/http"
)

var errNilContext = errors.New("context must be non-nil")

type Provider interface {
	//ServerInfo() bool
	StationProvider
	DepartureProvider
}

type Service struct {
	client *Client
}

type ClientOpt func(*Client)

type Client struct {
	userAgent  string
	httpClient *http.Client
	provider   Provider
	Station    *StationService
	Departure  *DepartureService
	debug      bool
}

func NewClient(provider Provider, opts ...ClientOpt) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
		provider:   provider,
		debug:      false,
	}
	c.initialize()

	for _, opt := range opts {
		opt(c)
	}

	if c.debug {
		c.httpClient.Transport = &util.LoggingTransport{}
	}

	return c
}

func (c *Client) initialize() {
	c.Station = &StationService{
		client: c,
	}
	c.Departure = &DepartureService{
		client: c,
	}
}

func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", ctx.Value("User-Agent").(string))

	resp, err := c.httpClient.Do(req)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, fmt.Errorf("failed to make station request: %w", err)
	}

	if resp.StatusCode < 200 && resp.StatusCode > 299 {
		return nil, fmt.Errorf("error from the api [statusCode:%d]", resp.StatusCode)
	}

	return resp, nil
}

func WithEnableDebugMode() ClientOpt {
	return func(c *Client) {
		c.debug = true
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOpt {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
