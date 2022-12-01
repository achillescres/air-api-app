package passlib

import (
	"crypto/sha256"
)

type HashManager interface {
	SHA256WithSalt(s string) (string, error)
}

type hashManager struct {
	salt string
}

var _ HashManager = (*hashManager)(nil)

func NewHashManager(salt string) HashManager {
	return &hashManager{salt: salt}
}

func (hM *hashManager) SHA256WithSalt(s string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return string(hash.Sum([]byte(hM.salt))), nil
}
