package ajwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type userClaims struct {
	jwt.StandardClaims
	User
}

type User struct {
	Id, Login, HashedPassword string
}

func newUserClaims(standardClaims jwt.StandardClaims, id, login, hashedPassword string) *userClaims {
	return &userClaims{StandardClaims: standardClaims,
		User: User{
			Id:             id,
			Login:          login,
			HashedPassword: hashedPassword,
		},
	}
}

func (m *jwtManager) User(id, login, hashedPassword string) (string, error) {
	now := time.Now()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, newUserClaims(
		jwt.StandardClaims{
			ExpiresAt: now.Add(m.jwtLiveTime).Unix(),
			IssuedAt:  now.Unix(),
			Subject:   "jwt user auth",
		},
		id,
		login,
		hashedPassword,
	))

	token, err := jwtToken.SignedString(m.secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *jwtManager) ParseUser(token string) (*User, error) {
	parsed, err := jwt.ParseWithClaims(token, userClaims{}, m.validateMethod)
	if err != nil {
		return nil, err
	}

	uClaims, ok := parsed.Claims.(userClaims)
	if !ok {
		return nil, errors.New("wrong claims type")
	}

	err = uClaims.Valid()
	if err != nil {
		return nil, err
	}

	return &uClaims.User, nil
}
