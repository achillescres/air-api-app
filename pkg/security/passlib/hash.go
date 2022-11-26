package passlib

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"sync"
)

var salt string
var inited = false
var once = &sync.Once{}

func Init(Salt string) {
	once.Do(func() {
		salt = Salt
		inited = true
	})
}

func Hash(s string) (string, error) {
	if !inited {
		log.Fatalln("need to init passlib to use it!")
		return "", errors.New("error you need to use Init before work")
	}

	hash := sha256.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		return "", err
	}
	sum := hash.Sum([]byte(salt))

	return fmt.Sprintf("%x", sum), nil
}
