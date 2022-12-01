package ajwt

import "time"

type Hasher interface {
	SHA256WithSalt(s string) (string, error)
}

type JWTManager interface {
	User(login, hashedPassword string) (string, error)
	RefreshToken() (string, int64, int64, error)
}

type jwtManager struct {
	hasher          Hasher
	secretKey       []byte
	jwtLiveTime     time.Duration
	refreshLiveTime time.Duration
}

func NewJWTManager(hasher Hasher, secretKey string, jwtLiveTime time.Duration, refreshLiveTime time.Duration) JWTManager {
	return &jwtManager{
		hasher:          hasher,
		secretKey:       []byte(secretKey),
		jwtLiveTime:     jwtLiveTime,
		refreshLiveTime: refreshLiveTime,
	}
}
