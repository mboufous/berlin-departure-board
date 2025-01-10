package internal

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"
)

type DepartureBoard struct {
	Station     Station     `json:"station"`
	Departures  []Departure `json:"departures"`
	LastUpdated time.Time   `json:"last_updated"`
}

type Departure struct {
	ID          string     `json:"id"`
	When        time.Time  `json:"when"`
	PlannedWhen *time.Time `json:"planned_when"`
	Delay       int        `json:"delay"`
	Product     Product    `json:"product"`
	Direction   string     `json:"direction"`
	Destination Station    `json:"destination"`
	Platform    string     `json:"platform"`
	Occupancy   string     `json:"occupancy"`
	Canceled    bool       `json:"canceled"`
	Hints       []Remark   `json:"hints"`
	Warnings    []Remark   `json:"warnings"`
}

type Product struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Remark struct {
	Type   string `json:"type"`
	Header string `json:"header"`
	Body   string `json:"body"`
}

type Station struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Lines []Product `json:"lines"` // Lines available at the station
}

type APIProviderStatus struct {
	Up bool `json:"up"`
}

type APIClient struct {
	registry             ProviderRegistry
	providers            []ApiProvider
	mu                   sync.Mutex
	index                int
	loadBalancing        bool
	selectedProviderName string
}

type ApiProvider interface {
	Name() string
	HealthCheck(ctx context.Context) (APIProviderStatus, error)
	GetDepartures(ctx context.Context, stationID string, config APIDeparturesParams) (DepartureBoard, error)
	GetStation(ctx context.Context, stationID string) (Station, error)
}

type Option func(client *APIClient)

func NewTransportAPIClient(registry ProviderRegistry, Opts ...Option) (*APIClient, error) {
	providers := registry.GetProviders()
	if len(providers) == 0 {
		return nil, fmt.Errorf("no providers available")
	}

	client := &APIClient{
		registry:  registry,
		providers: providers,
	}

	for _, opt := range Opts {
		opt(client)
	}

	// Validation to prevent both load balancing and selected provider
	if err := client.validate(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *APIClient) validate() error {
	if c.loadBalancing && len(c.selectedProviderName) != 0 {
		return fmt.Errorf("cannot enable load balancing and select a specific provider at the same time")
	}
	return nil
}

func WithLoadBalancer() Option {
	return func(client *APIClient) {
		if len(client.providers) == 1 {
			slog.Info("Load balancing disabled, only one provider available")
			return
		}
		slog.Info("Load balancing enabled")
		client.loadBalancing = true
	}
}

func WithSelectedProvider(providerName string) Option {
	return func(client *APIClient) {
		slog.Info("Selected provider", "provider", providerName)
		client.selectedProviderName = providerName
	}
}

func (c *APIClient) nextProvider() ApiProvider {
	c.mu.Lock()
	provider := c.providers[c.index]
	c.index = (c.index + 1) % len(c.providers)
	c.mu.Unlock()
	slog.Info("Next provider",
		"provider_index", c.index,
		"provider", provider,
	)
	return provider
}

func (c *APIClient) selectProvider() (ApiProvider, error) {
	if c.loadBalancing {
		return c.nextProvider(), nil
	}

	if c.selectedProviderName != "" {
		if provider := c.registry.GetProvider(c.selectedProviderName); provider != nil {
			return provider, nil
		}
		return nil, fmt.Errorf("selected provider %s not found", c.selectedProviderName)
	}

	return c.providers[0], nil
}

func (c *APIClient) GetDepartures(ctx context.Context, stationID string, config APIDeparturesParams) (DepartureBoard, error) {
	provider, err := c.selectProvider()
	if err != nil {
		return DepartureBoard{}, err
	}
	return provider.GetDepartures(ctx, stationID, config)
}

func (c *APIClient) GetStation(ctx context.Context, stationID string) (Station, error) {
	provider, err := c.selectProvider()
	if err != nil {
		return Station{}, err
	}
	return provider.GetStation(ctx, stationID)
}

func (c *APIClient) HealthCheck(ctx context.Context) (APIProviderStatus, error) {
	provider, err := c.selectProvider()
	if err != nil {
		return APIProviderStatus{}, err
	}
	return provider.HealthCheck(ctx)
}
