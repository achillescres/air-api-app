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
	secret []byte
	inited = false
	once   = &sync.Once{}
)

func Init(Secret string) {
	once.Do(func() {
		inited = true
		secret = []byte(Secret)
	})
}

type UserClaims struct {
	*jwt.StandardClaims

	Login          string
	HashedPassword string
}

func newUserClaims(standardClaims *jwt.StandardClaims, login string, hashedPassword string) *UserClaims {
	return &UserClaims{StandardClaims: standardClaims, Login: login, HashedPassword: hashedPassword}
}

func User(login, hashedPassword string) (string, error) {
	if !inited {
		log.Fatalln("you need to init ajwt to use it!")
	}

	now := time.Now()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, newUserClaims(
		&jwt.StandardClaims{
			ExpiresAt: now.Add(liveTime).Unix(),
			IssuedAt:  now.Unix(),
			Subject:   "user auth",
		},
		login,
		hashedPassword,
	))

	token, err := jwtToken.SignedString(secret)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func ValidateUser(token string) error {
	parse, err := jwt.Parse(token)
	if err != nil {
		return err
	}
}
