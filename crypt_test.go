package gofacades

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypt(t *testing.T) {
	assert := assert.New(t)
	secretKey := "Om8FLaOZc0Y2IVx58K9MGTgm8RCmmE0L"
	str := "123456"

	t.Run("encrypt string", func(t *testing.T) {
		_, err := Crypt(secretKey).EncryptString(str)

		assert.NoError(err)
	})

	t.Run("decrypt string", func(t *testing.T) {
		encrypted, err := Crypt(secretKey).EncryptString(str)

		// Assert encryption error
		assert.NoError(err)

		// Decrypt string
		decrypted, err := Crypt(secretKey).DecryptString(encrypted)

		assert.NoError(err)
		assert.Equal(str, decrypted)
	})
}
