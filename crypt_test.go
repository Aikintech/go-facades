package gofacades

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypt(t *testing.T) {
	assert := assert.New(t)
	secretKey := "Om8FLaOZc0Y2IVx58K9MGTgm8RCmmE0L"
	str := "123456"

	t.Run("invalid secret key", func(t *testing.T) {
		crypt := Crypt("1234")

		assert.Empty(crypt)
	})

	t.Run("encrypt string", func(t *testing.T) {
		_, err := Crypt(secretKey).EncryptString(str)

		assert.NoError(err)
	})

	t.Run("decrypt string", func(t *testing.T) {
		crypt := Crypt(secretKey)
		encrypted, err := crypt.EncryptString(str)

		// Assert encryption error
		assert.NoError(err)

		decrypted, err := crypt.DecryptString(encrypted)

		assert.NoError(err)
		assert.Equal(str, decrypted)
	})
}
