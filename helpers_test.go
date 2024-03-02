package gofacades

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileExtension(t *testing.T) {
	assert := assert.New(t)

	t.Run("returns an empty extension", func(t *testing.T) {
		mime := ""
		ext := GetFileExtension(mime)

		assert.Equal("", ext)
	})

	t.Run("returns the correct extension", func(t *testing.T) {
		mime := "image/jpg"
		ext := GetFileExtension(mime)

		assert.Equal("jpg", ext)
	})
}

func TestGenerateRandomString(t *testing.T) {
	assert := assert.New(t)

	t.Run("returns an empty string", func(t *testing.T) {
		result := GenerateRandomString(0, false)

		assert.Equal("", result)
	})

	t.Run("returns the correct number of characters", func(t *testing.T) {
		length := 20
		result := GenerateRandomString(uint(length), false)

		assert.Equal(length, len(result))
	})
}

func TestGenerateRandomNumbers(t *testing.T) {
	assert := assert.New(t)

	t.Run("returns zero", func(t *testing.T) {
		length := 0
		result := GenerateRandomNumbers(uint(length))

		assert.Equal(int64(0), result)
	})

	t.Run("returns the correct number of characters", func(t *testing.T) {
		length := 10
		result := GenerateRandomNumbers(uint(length))
		str := strconv.FormatInt(result, 10)

		assert.NotEqual(int64(0), result)
		assert.Equal(length, len(str))
	})
}
