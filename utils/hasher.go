package utils

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Generate(password string) ([]byte, error)
	Compare(password string, hash string) error
}

type BcryptHasher struct {
}

func NewBcryptHasher() Hasher {
	return &BcryptHasher{}
}

func (b *BcryptHasher) Generate(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}

func (b *BcryptHasher) Compare(h string, p string) error {
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
}
