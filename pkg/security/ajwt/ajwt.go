package ajwt

import (
	"github.com/golang-jwt/jwt"
	"log"
	"sync"
	"time"
)

const (
	liveTime = 30 * time.Minute
)

var (
	secret = ""
	inited = false
	once   = &sync.Once{}
)

func Init(Secret string) {
	once.Do(func() {
		inited = true
		secret = Secret
	})
}

type UserClaims struct {
	jwt.StandardClaims
	Id             string
	Login          string
	HashedPassword string
}

func User(login, hashedPassword string) (string, error) {
	if !inited {
		log.Fatalln("you need to init ajwt to use it!")
	}

	now := time.Now()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(liveTime).Unix(),
			IssuedAt:  now.Unix(),
			Subject:   "user auth",
		},
		Login:          login,
		HashedPassword: hashedPassword,
	})
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", nil
	}

	return token, nil
}
