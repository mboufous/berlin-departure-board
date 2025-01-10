package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock implementations for testing
type MockProvider struct {
	mock.Mock
	name string
}

func (p *MockProvider) Name() string {
	return p.name
}

func (p *MockProvider) GetDepartures(ctx context.Context, stationID string) (DepartureBoard, error) {
	args := p.Called(ctx, stationID)
	return args.Get(0).([]Departure), args.Error(1)
}

func (p *MockProvider) HealthCheck() (APIProviderStatus, error) {
	return APIProviderStatus{Status: "up"}, nil
}

type MockProviderRegistry struct {
	mock.Mock
}

func (r *MockProviderRegistry) GetProviders() []ApiProvider {
	args := r.Called()
	return args.Get(0).([]ApiProvider)
}

func (r *MockProviderRegistry) GetProvider(name string) ApiProvider {
	args := r.Called(name)

	// Handle nil correctly by ensuring the return type is explicitly ApiProvider
	if provider, ok := args.Get(0).(ApiProvider); ok {
		return provider
	}
	return nil
}
func TestNewTransportAPIClient_NoProviders(t *testing.T) {
	registry := &MockProviderRegistry{}
	registry.On("GetProviders").Return([]ApiProvider{})

	client, err := NewTransportAPIClient(registry)
	assert.Nil(t, client)
	assert.EqualError(t, err, "no providers available")
}

func TestNewTransportAPIClient_OneProvider(t *testing.T) {
	provider := &MockProvider{name: "ProviderA"}
	registry := &MockProviderRegistry{}
	registry.On("GetProviders").Return([]ApiProvider{provider})

	client, err := NewTransportAPIClient(registry)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, 1, len(client.providers))
}

func TestNewTransportAPIClient_MultipleProviders(t *testing.T) {
	providerA := &MockProvider{name: "ProviderA"}
	providerB := &MockProvider{name: "ProviderB"}
	registry := &MockProviderRegistry{}
	registry.On("GetProviders").Return([]ApiProvider{providerA, providerB})

	client, err := NewTransportAPIClient(registry, WithLoadBalancer())
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, 2, len(client.providers))
	assert.True(t, client.loadBalancing)
}

func TestNewTransportAPIClient_WithSelectedProvider(t *testing.T) {
	providerA := &MockProvider{name: "ProviderA"}
	registry := &MockProviderRegistry{}
	registry.On("GetProviders").Return([]ApiProvider{providerA})
	registry.On("GetProvider", "ProviderA").Return(providerA)

	client, err := NewTransportAPIClient(registry, WithSelectedProvider("ProviderA"))
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "ProviderA", client.selectedProviderName)

	provider, err := client.selectProvider()
	assert.NoError(t, err)
	assert.Equal(t, providerA, provider)
}

func TestNewTransportAPIClient_WithLoadBalancerAndSelectedProvider(t *testing.T) {
	providerA := &MockProvider{name: "ProviderA"}
	providerB := &MockProvider{name: "ProviderB"}
	registry := &MockProviderRegistry{}
	registry.On("GetProviders").Return([]ApiProvider{providerA, providerB})

	client, err := NewTransportAPIClient(registry, WithLoadBalancer(), WithSelectedProvider("ProviderA"))
	assert.Nil(t, client)
	assert.EqualError(t, err, "cannot enable load balancing and select a specific provider at the same time")
}

func TestAPIClient_GetDepartures_LoadBalancing(t *testing.T) {
	providerA := &MockProvider{name: "ProviderA"}
	providerB := &MockProvider{name: "ProviderB"}
	registry := &MockProviderRegistry{}
	registry.On("GetProviders").Return([]ApiProvider{providerA, providerB})

	client, err := NewTransportAPIClient(registry, WithLoadBalancer())
	assert.NoError(t, err)

	providerA.On("GetDepartures", mock.Anything, "stop123").Return([]Departure{{Line: "A1"}}, nil)
	providerB.On("GetDepartures", mock.Anything, "stop123").Return([]Departure{{Line: "B1"}}, nil)

	departures, err := client.GetDepartures(context.Background(), "stop123")
	assert.NoError(t, err)
	assert.Equal(t, "A1", departures[0].Line)

	departures, err = client.GetDepartures(context.Background(), "stop123")
	assert.NoError(t, err)
	assert.Equal(t, "B1", departures[0].Line)
}

func TestAPIClient_GetDepartures_SelectedProvider(t *testing.T) {
	providerA := &MockProvider{name: "ProviderA"}
	registry := &MockProviderRegistry{}
	registry.On("GetProviders").Return([]ApiProvider{providerA})
	registry.On("GetProvider", "ProviderA").Return(providerA)

	client, err := NewTransportAPIClient(registry, WithSelectedProvider("ProviderA"))
	assert.NoError(t, err)

	providerA.On("GetDepartures", mock.Anything, "stop123").Return([]Departure{{Line: "A1"}}, nil)

	departures, err := client.GetDepartures(context.Background(), "stop123")
	assert.NoError(t, err)
	assert.Equal(t, "A1", departures[0].Line)
}

func TestAPIClient_GetDepartures_NoProviderFound(t *testing.T) {
	providerA := &MockProvider{name: "ProviderA"}
	registry := &MockProviderRegistry{}
	registry.On("GetProviders").Return([]ApiProvider{providerA})

	// Ensure the mock returns nil explicitly as ApiProvider
	registry.On("GetProvider", "NonExistent").Return(nil).Run(func(args mock.Arguments) {
		return // Explicitly type nil for ApiProvider
	})

	client, err := NewTransportAPIClient(registry, WithSelectedProvider("NonExistent"))
	assert.NoError(t, err)

	_, err = client.GetDepartures(context.Background(), "stop123")
	assert.EqualError(t, err, "selected provider NonExistent not found")
}
