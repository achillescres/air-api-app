package ajwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type userClaims struct {
	*jwt.StandardClaims

	Login          string
	HashedPassword string
}

func newUserClaims(standardClaims *jwt.StandardClaims, login string, hashedPassword string) *userClaims {
	return &userClaims{StandardClaims: standardClaims, Login: login, HashedPassword: hashedPassword}
}

func (m *jwtManager) User(login, hashedPassword string) (string, error) {
	now := time.Now()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, newUserClaims(
		&jwt.StandardClaims{
			ExpiresAt: now.Add(m.jwtLiveTime).Unix(),
			IssuedAt:  now.Unix(),
			Subject:   "jwt user auth",
		},
		login,
		hashedPassword,
	))

	token, err := jwtToken.SignedString(m.secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

//func ValidateUser(token string) error {
//	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
//
//	})
//	if err != nil {
//		return err
//	}
//}
