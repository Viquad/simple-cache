package simplecache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	errMissedValue = errors.New("missed value")
)

type Cache interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type memoryCache struct {
	impl map[string]*entity
	mu   sync.Mutex
}

type entity struct {
	deadline *time.Time
	value    interface{}
}

func (e entity) isExpired() bool {
	if e.deadline != nil {
		return e.deadline.Before(time.Now())
	}

	return false
}

func NewMemoryCache() *memoryCache {
	return &memoryCache{
		impl: make(map[string]*entity),
		mu:   sync.Mutex{},
	}
}

func (c *memoryCache) Set(key string, value interface{}, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	e := entity{
		value: value,
	}

	if ttl > 0 {
		deadline := time.Now().Add(ttl)
		e.deadline = &deadline
		go func() {
			time.Sleep(ttl)
			c.mu.Lock()
			if e, ok := c.impl[key]; ok {
				if e.isExpired() {
					delete(c.impl, key)
				}
			}
			c.mu.Unlock()
		}()
	}

	c.impl[key] = &e

	return nil
}

func (c *memoryCache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.impl[key]
	if !ok {
		return nil, fmt.Errorf("%w by %s key", errMissedValue, key)
	}

	return e.value, nil
}

func (c *memoryCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.impl[key]; !ok {
		return fmt.Errorf("attempt to delete %w by %s key", errMissedValue, key)
	}

	delete(c.impl, key)

	return nil
}
