package cache

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrObjectNotFound = errors.New("object not found in cache store")
)

type Store[T any] interface {
	Get(ctx context.Context, key string) (T, error)
	GetWithTTL(ctx context.Context, key string) (T, time.Duration, error)
	Put(ctx context.Context, key string, object T, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Clear(ctx context.Context) error
}
type Encoder interface {
	Encode(any) ([]byte, error)
	Decode(data []byte, returnObject any) error
}

type Cache struct {
	store   Store[[]byte]
	encoder Encoder
}

func NewCache(store Store[[]byte], encoder Encoder) *Cache {
	return &Cache{
		store:   store,
		encoder: encoder,
	}
}

func (c *Cache) Get(ctx context.Context, key string, returnObject any) error {
	rawCachedObject, err := c.store.Get(ctx, key)
	if err != nil {
		return err
	}

	if err := c.encoder.Decode(rawCachedObject, returnObject); err != nil {
		return fmt.Errorf("error decoding station: %w", err)
	}

	return nil
}

func (c *Cache) Put(ctx context.Context, key string, object any, ttl time.Duration) error {
	encodedObject, err := c.encoder.Encode(object)
	if err != nil {
		return fmt.Errorf("couldn't encode station: %w", err)
	}

	return c.store.Put(ctx, key, encodedObject, ttl)
}

//func (c *Cache) GetWithTTL(ctx context.Context, key string) (T, time.Duration, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (c *Cache) Delete(ctx context.Context, key string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (c *Cache) Clear(ctx context.Context) error {
//	//TODO implement me
//	panic("implement me")
//}
