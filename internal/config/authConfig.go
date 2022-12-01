package config

import (
	"path"
	"sync"
	"time"
)

const authConfigFilename = "auth.config.yaml"

type AuthConfig struct {
	JWTLiveTimeMinutes          int64 `yaml:"JWTLiveTimeMinutes" env-required:"true"`
	RefreshTokenLiveTimeMinutes int64 `yaml:"refreshTokenLiveTimeMinutes" env-required:"true"`

	JWTLiveTime          time.Duration
	RefreshTokenLiveTime time.Duration
}

var (
	authConfigInst     = &AuthConfig{}
	loadAuthConfigOnce = sync.Once{}
)

func Auth() AuthConfig {
	loadAuthConfigOnce.Do(func() {
		readConfig(path.Join(Env().ConfigAbsPath, authConfigFilename), authConfigInst)
		authConfigInst.JWTLiveTime = time.Minute * time.Duration(authConfigInst.JWTLiveTimeMinutes)
		authConfigInst.RefreshTokenLiveTime = time.Minute * time.Duration(authConfigInst.RefreshTokenLiveTimeMinutes)
	})

	return *authConfigInst
}
