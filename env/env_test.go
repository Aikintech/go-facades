package gofacades

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	assert := assert.New(t)

	t.Run("env get string", func(t *testing.T) {
		var (
			key   = "str"
			value = "string"
		)

		// Set string value
		err := os.Setenv(key, value)
		assert.NoError(err)

		// Get value from env
		val := Env().GetString(key)

		assert.Equal(value, val)
	})

	t.Run("env get bool", func(t *testing.T) {
		var (
			key   = "str"
			value = "true"
		)

		// Set string value
		err := os.Setenv(key, value)
		assert.NoError(err)

		// Get value from env
		val := Env().GetBool(key)

		assert.Equal(true, val)
	})

	t.Run("env get int", func(t *testing.T) {
		var (
			key   = "str"
			value = "1"
		)

		// Set string value
		err := os.Setenv(key, value)
		assert.NoError(err)

		// Get value from env
		val := Env().GetInt(key)

		assert.Equal(1, val)
	})

	t.Run("env get slice", func(t *testing.T) {
		var (
			key   = "str"
			value = "1,2"
		)

		// Set string value
		err := os.Setenv(key, value)
		assert.NoError(err)

		// Get value from env
		val := Env().GetStringSlice(key)

		assert.Equal([]string{"1", "2"}, val)
	})
}
