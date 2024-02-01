package hafas

import (
	"context"
	"errors"
	"fmt"
	"github.com/mboufous/berlin-departure-board/util"
	"io"
	"net/http"
)

var errNilContext = errors.New("context must be non-nil")

type StationProvider interface {
	NewStationRequest(params any) (*http.Request, error)
	NewStationResponse(body io.ReadCloser) (*Station, error)
}

type DepartureProvider interface {
	NewDepartureRequest(param any) (*http.Request, error)
	NewDepartureResponse(body io.ReadCloser) (*Departure, error)
}

type Provider interface {
	//ServerInfo() bool
	StationProvider
	DepartureProvider
}

type Service struct {
	client *Client
}

type Client struct {
	userAgent  string
	httpClient *http.Client
	provider   Provider
	Station    *StationService
	debug      bool
}
type ClientOpt func(*Client)

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
}

func (c *Client) NewStationRequest(params any) (*http.Request, error) {
	return c.provider.NewStationRequest(params)
}

func (c *Client) NewDepartureRequest(params any) (*http.Request, error) {
	return c.provider.NewDepartureRequest(param)
}

func (c *Client) Do(ctx context.Context, req *http.Request) (*Station, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, fmt.Errorf("failed to make station request: %w", err)
	}

	if resp.StatusCode < 200 && resp.StatusCode > 299 {
		// TODO: unmarshal the error we got from the api and encapsulate it to our error
		return nil, errors.New(fmt.Sprintf("error from the api [statusCode:%d]", resp.StatusCode))
	}

	return c.provider.NewStationResponse(resp.Body)
}

type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Message  string         `json:"message"`
}

func (r *ErrorResponse) Error() string {
	if r.Response != nil && r.Response.Request != nil {
		return fmt.Sprintf("%v %v: %d %v",
			r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
	}

	if r.Response != nil {
		return fmt.Sprintf("%d %v", r.Response.StatusCode, r.Message)
	}

	return fmt.Sprintf("%v", r.Message)
}

//func (c *HafasClient) Ping() {
//	c.provider.ServerInfo()
//}

//func (c *HafasClient) GetStation(ctx context.Context, id string) (string, error) {
//	// create bvg http request to get the station
//	payload := c.provider.CreateStationRequestPayload(id)
//	jsonRequest, err := json.Marshal(payload)
//	if err != nil {
//		return "", err
//	}
//	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(jsonRequest))
//	if err != nil {
//		return "", err
//	}
//	resp, err := c.httpClient.Do(request)
//	if err != nil {
//		return "", err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return "", errors.New("station http request failed")
//	}
//
//	return
//}

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
