package gofacades

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2 struct {
	format  string
	version int
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
	saltLen uint32
}

func NewArgon2() Hash {
	return &Argon2{
		format:  "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		version: argon2.Version,
		time:    uint32(4),
		memory:  uint32(65536),
		threads: uint8(1),
		keyLen:  32,
		saltLen: 16,
	}
}

func (a *Argon2) Make(value string) (string, error) {
	salt := make([]byte, a.saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(value), salt, a.time, a.memory, a.threads, a.keyLen)

	return fmt.Sprintf(a.format, a.version, a.memory, a.time, a.threads, base64.RawStdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(hash)), nil
}

func (a *Argon2) Check(value, hash string) bool {
	hashParts := strings.Split(hash, "$")
	if len(hashParts) != 6 {
		return false
	}

	var version int
	_, err := fmt.Sscanf(hashParts[2], "v=%d", &version)
	if err != nil {
		return false
	}
	if version != a.version {
		return false
	}

	memory := a.memory
	time := a.time
	threads := a.threads

	_, err = fmt.Sscanf(hashParts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(hashParts[4])
	if err != nil {
		return false
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(hashParts[5])
	if err != nil {
		return false
	}

	hashToCompare := argon2.IDKey([]byte(value), salt, time, memory, threads, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1
}

func (a *Argon2) NeedsRehash(hash string) bool {
	hashParts := strings.Split(hash, "$")
	if len(hashParts) != 6 {
		return true
	}

	var version int
	_, err := fmt.Sscanf(hashParts[2], "v=%d", &version)
	if err != nil {
		return true
	}
	if version != a.version {
		return true
	}

	var memory, time uint32
	var threads uint8
	_, err = fmt.Sscanf(hashParts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return true
	}

	return memory != a.memory || time != a.time || threads != a.threads
}
