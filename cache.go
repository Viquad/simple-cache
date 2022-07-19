package simplecache

import (
	"errors"
	"fmt"
)

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type memoryCache struct {
	impl map[string]interface{}
}

func NewMemoryCache() *memoryCache {
	return &memoryCache{impl: make(map[string]interface{})}
}

func (c *memoryCache) Set(key string, value interface{}) error {
	if len(key) < 1 {
		return errors.New("empty key")
	}

	c.impl[key] = value
	return nil
}

func (c *memoryCache) Get(key string) (interface{}, error) {
	if len(key) < 1 {
		return nil, errors.New("empty key")
	}

	value, ok := c.impl[key]
	if !ok {
		return nil, fmt.Errorf("missed value by %s key in memory cache", key)
	}

	return value, nil
}

func (c *memoryCache) Delete(key string) error {
	if len(key) < 1 {
		return errors.New("empty key")
	}

	_, ok := c.impl[key]
	if !ok {
		return fmt.Errorf("attempt to delete missed value by %s key in memory cache", key)
	}

	delete(c.impl, key)

	return nil
}

/*
func main() {
	cache := NewMemoryCache()

	cache.Set("userId", 42)
	userId, err := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}
*/
