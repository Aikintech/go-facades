package gofacades

import (
	"fmt"
	"github.com/gofiber/storage/redis/v3"
	"time"
)

type RedisCache struct {
	client *redis.Storage
	prefix string
}

var cache RedisCache

func NewConnection(config redis.Config, prefix string) {
	cache = RedisCache{
		client: redis.New(config),
		prefix: prefix,
	}
}

func Redis() *RedisCache {
	return &RedisCache{
		client: cache.client,
		prefix: cache.prefix,
	}
}

// Delete removes an item from the cache.
func (r *RedisCache) Delete(key string) error {
	return r.client.Delete(r.key(key))
}

// Flush removes all items from the cache.
func (r *RedisCache) Flush() error {
	return r.client.Reset()
}

// Forever puts an item in the cache indefinitely.
func (r *RedisCache) Forever(key string, value []byte) error {
	return r.Put(key, value, 0)
}

// Forget deletes an item from the cache.
func (r *RedisCache) Forget(key string) error {
	return r.Delete(key)
}

// Get retrieves an item by key
func (r *RedisCache) Get(key string) ([]byte, error) {
	return r.client.Get(r.key(key))
}

// Set sets an item in the cache for a given time.
func (r *RedisCache) Set(key string, value []byte, expiration time.Duration) error {
	return r.Put(key, value, expiration)
}

// Put puts an item in the cache for a given time.
func (r *RedisCache) Put(key string, value []byte, expiration time.Duration) error {
	return r.client.Set(r.key(key), value, expiration)
}

// Remember gets an item from the cache, or execute the given Closure and store the result for a given time.
func (r *RedisCache) Remember(key string, expiration time.Duration, callback func() ([]byte, error)) ([]byte, error) {
	data, err := r.Get(key)

	if err != nil {
		// Invoke the callback
		data, err = callback()

		if err != nil {
			return nil, err
		}

		// Store the value in the cache
		err = r.Put(key, data, expiration)

		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// RememberForever gets an item from the cache, or execute the given Closure and store the result indefinitely.
func (r *RedisCache) RememberForever(key string, callback func() ([]byte, error)) ([]byte, error) {
	data, err := r.Get(key)

	if err != nil {
		// Invoke the callback
		data, err = callback()

		if err != nil {
			return nil, err
		}

		// Store the value in the cache
		err = r.Put(key, data, 0)

		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// format key with prefix
func (r *RedisCache) key(key string) string {
	var formattedKey string

	if len(r.prefix) > 0 {
		formattedKey = fmt.Sprintf("%s:%s", r.prefix, key)
	}

	return formattedKey
}
