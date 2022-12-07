package ajwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type Hasher interface {
	SHA256WithSalt(s string) (string, error)
}

type JWTManager interface {
	User(id, login, hashedPassword string) (string, error)
	ParseUser(token string) (*User, error)
	RefreshToken() (string, int64, int64, error)
	validateMethod(token *jwt.Token) (any, error)
}

type jwtManager struct {
	hasher          Hasher
	secretKey       []byte
	jwtLiveTime     time.Duration
	refreshLiveTime time.Duration
	keyFuncFabric   func(secret []byte) jwt.Keyfunc
}

func NewJWTManager(hasher Hasher, secretKey string, jwtLiveTime time.Duration, refreshLiveTime time.Duration) JWTManager {
	return &jwtManager{
		hasher:          hasher,
		secretKey:       []byte(secretKey),
		jwtLiveTime:     jwtLiveTime,
		refreshLiveTime: refreshLiveTime,
	}
}

func (m *jwtManager) validateMethod(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("wrong signing method")
	}

	return m.secretKey, nil
}
