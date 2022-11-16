package passlib

import (
	"errors"
	"fmt"
	"sync"
)

// TODO implement passlib

var salt string
var hahser any
var inited = false
var once = &sync.Once{}

//func Init(Salt string, hashMethod method) {
//	once.Do(func() {
//		salt = Salt
//		hasher.Set(hashMethod)
//		inited = true
//	})
//}

func HashPassword(password string) (string, error) {
	inited = true // TODO DELETE THIS
	if !inited {
		return "", errors.New("error you need to passlib.Init to work")
	}

	return fmt.Sprintf("hashof:%s", password), nil
}
