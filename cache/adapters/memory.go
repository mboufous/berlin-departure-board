package adapters

import (
	"context"
	"github.com/mboufous/berlin-departure-board/cache"
	gocache "github.com/patrickmn/go-cache"
	"time"
)

type MemoryAdapter[T any] struct {
	provider *gocache.Cache
}

func NewMemoryAdapter[T any](defaultExpiration, cleanupInterval time.Duration) *MemoryAdapter[T] {
	return &MemoryAdapter[T]{
		provider: gocache.New(defaultExpiration, cleanupInterval),
	}
}

func (s *MemoryAdapter[T]) Get(_ context.Context, key string) (T, error) {
	if v, ok := s.provider.Get(key); ok {
		return v.(T), nil
	}
	var emptyObj T
	return emptyObj, cache.ErrObjectNotFound
}

func (s *MemoryAdapter[T]) GetWithTTL(_ context.Context, key string) (T, time.Duration, error) {
	if v, ttl, ok := s.provider.GetWithExpiration(key); ok {
		return v.(T), time.Until(ttl), nil
	}
	var emptyObj T
	return emptyObj, 0, cache.ErrObjectNotFound
}

func (s *MemoryAdapter[T]) Put(_ context.Context, key string, object T, ttl time.Duration) error {
	s.provider.Set(key, object, ttl)
	return nil
}

func (s *MemoryAdapter[T]) Delete(_ context.Context, key string) error {
	s.provider.Delete(key)
	return nil
}

func (s *MemoryAdapter[T]) Clear(_ context.Context) error {
	s.provider.Flush()
	return nil
}
