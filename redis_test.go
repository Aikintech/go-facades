package gofacades

import (
	"testing"
	"time"

	"github.com/gofiber/storage/redis/v3"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	var prefix = "test"
	var key = "redis"
	var value = "12345"
	var assert = assert.New(t)

	t.Run("redis set/put", func(t *testing.T) {
		rdb := Redis(redis.Config{Reset: true}, prefix)

		err := rdb.Set(key, []byte(value), time.Second)

		assert.NoError(err)
	})

	t.Run("redis expired", func(t *testing.T) {
		rdb := Redis(redis.Config{Reset: true}, prefix)

		err := rdb.Set(key, []byte(value), time.Millisecond*200)
		assert.NoError(err)

		time.Sleep(time.Second)

		val, err := rdb.Get(key)

		assert.NoError(err)
		assert.Equal("", string(val))
	})

	t.Run("redis delete/forget", func(t *testing.T) {
		rdb := Redis(redis.Config{Reset: true}, prefix)
		err := rdb.Set(key, []byte(value), time.Minute)
		assert.NoError(err)

		err = rdb.Delete(key)
		assert.NoError(err)
	})

	t.Run("redis get", func(t *testing.T) {
		rdb := Redis(redis.Config{Reset: true}, prefix)
		err := rdb.Set(key, []byte(value), time.Minute)
		assert.NoError(err)

		val, err := rdb.Get(key)

		assert.NoError(err)
		assert.Equal([]byte(value), val)
	})

	t.Run("redis flush", func(t *testing.T) {
		rdb := Redis(redis.Config{Reset: true}, prefix)
		err := rdb.Set(key, []byte(value), time.Minute)

		assert.NoError(err)

		err = rdb.Flush()

		assert.NoError(err)
	})

	t.Run("redis forever", func(t *testing.T) {
		rdb := Redis(redis.Config{Reset: true}, prefix)
		err := rdb.Forever(key, []byte(value))

		assert.NoError(err)

		time.Sleep(time.Second * 2)

		val, err := rdb.Get(key)

		assert.NoError(err)
		assert.Equal([]byte(value), val)
	})
}
