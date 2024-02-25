package gofacades

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"github.com/gookit/color"
)

type AesCrypto struct {
	key []byte
}

func Crypt(secretKey string) *AesCrypto {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 && len(secretKey) != 64 {
		color.Redln("Empty or invalid secret key provided. Key must be 16 or 24 or 32 or 64 characters")
		return nil
	}
	keyBytes := []byte(secretKey)
	return &AesCrypto{
		key: keyBytes,
	}
}

// EncryptString encrypts the given string, and returns the iv and cipher text as base64 encoded strings.
func (b *AesCrypto) EncryptString(value string) (string, error) {
	block, err := aes.NewCipher(b.key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(value)

	iv := make([]byte, 12)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesGcm.Seal(nil, iv, plaintext, nil)

	var jsonEncoded []byte
	jsonEncoded, err = json.Marshal(map[string][]byte{
		"iv":    iv,
		"value": cipherText,
	})

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(jsonEncoded), nil
}

// DecryptString decrypts the given iv and cipher text, and returns the plaintext.
func (b *AesCrypto) DecryptString(payload string) (string, error) {
	decodePayload, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return "", err
	}

	decodeJson := make(map[string][]byte)
	err = json.Unmarshal(decodePayload, &decodeJson)
	if err != nil {
		return "", err
	}

	// check if the json payload has the correct keys
	if _, ok := decodeJson["iv"]; !ok {
		return "", errors.New("decrypt payload error: missing iv key")
	}
	if _, ok := decodeJson["value"]; !ok {
		return "", errors.New("decrypt payload error: missing value key")
	}

	decodeIv := decodeJson["iv"]
	decodeCipherText := decodeJson["value"]

	block, err := aes.NewCipher(b.key)
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesGcm.Open(nil, decodeIv, decodeCipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
