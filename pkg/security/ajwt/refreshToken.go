package ajwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type refreshTokenClaims struct {
	*jwt.StandardClaims
	Salt string
}

func newRefreshTokenClaims(standardClaims *jwt.StandardClaims, salt string) *refreshTokenClaims {
	return &refreshTokenClaims{StandardClaims: standardClaims, Salt: salt}
}

func (m *jwtManager) RefreshToken() (string, int64, int64, error) {
	salt, err := m.hasher.SHA256WithSalt(uuid.New().String())
	if err != nil {
		return "", -1, -1, err
	}

	issuedTime, expiresTime := time.Now().Unix(), time.Now().Add(m.refreshLiveTime).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, newRefreshTokenClaims(
		&jwt.StandardClaims{
			IssuedAt:  issuedTime,
			ExpiresAt: expiresTime,
			Subject:   "rt user auth",
		},
		salt,
	))

	token, err := jwtToken.SignedString(m.secretKey)
	if err != nil {
		return "", -1, -1, nil
	}

	return token, issuedTime, expiresTime, nil
}
