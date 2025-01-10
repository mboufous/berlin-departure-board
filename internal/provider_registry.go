package internal

import (
	"context"
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"sync"
	"time"
)

const (
	providerHealthCheckTimout = 5 * time.Second
)

type ProviderRegistry interface {
	GetProviders() []ApiProvider
	GetProvider(name string) ApiProvider
}

type DefaultProviderRegistry struct {
	providers map[string]ApiProvider
	mu        sync.Mutex
}

func NewProviderRegistry(providers ...ApiProvider) (*DefaultProviderRegistry, error) {
	r := &DefaultProviderRegistry{
		providers: make(map[string]ApiProvider, len(providers)),
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, p := range providers {
		ctx, cancel := context.WithTimeout(context.Background(), providerHealthCheckTimout)
		defer cancel()
		status, err := p.HealthCheck(ctx)
		if err != nil || !status.Up {
			return nil, fmt.Errorf("provider is not healthy: %w", err)
		}
		r.providers[p.Name()] = p
		slog.Info("Added Healthy Provider", "Provider", p.Name())
	}
	slog.Info("Providers registered", "count", len(r.providers))

	return r, nil
}

func (r *DefaultProviderRegistry) GetProvider(name string) ApiProvider {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.providers[name]
}

func (r *DefaultProviderRegistry) GetProviders() []ApiProvider {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Order is not important
	return slices.Collect(maps.Values(r.providers))
}

// Run inside a goroutine
//func (r *DefaultProviderRegistry) periodicHealthCheck() {
//	ticker := time.NewTicker(1 * time.Minute)
//	defer ticker.Stop()
//
//	for range ticker.C {
//		r.mu.Lock()
//		var healthyProviders []APIClient
//		for _, provider := range r.providers {
//			status, err := provider.HealthCheck()
//			if err != nil || status.Status != "up" {
//				r.log.Warnf("Provider unhealthy: %s, reason: %s", status.Provider, status.Message)
//				continue
//			}
//			healthyProviders = append(healthyProviders, provider)
//		}
//		r.providers = healthyProviders
//		r.mu.Unlock()
//	}
//}
