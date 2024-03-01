package gofacades

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	rounds int
}

func NewBcrypt() Hash {
	return &Bcrypt{
		rounds: 10,
	}
}

func (b *Bcrypt) Make(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), b.rounds)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (b *Bcrypt) Check(value, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}

func (b *Bcrypt) NeedsRehash(hash string) bool {
	hashCost, err := bcrypt.Cost([]byte(hash))

	if err != nil {
		return true
	}
	return hashCost != b.rounds
}
