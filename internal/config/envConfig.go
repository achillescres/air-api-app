package config

import (
	"github.com/achillescres/saina-api/pkg/gconfig"
	log "github.com/sirupsen/logrus"
	"path"
	"sync"
)

const envFilename = ".env"

// EnvConfig Env
type EnvConfig struct {
	ProjectAbsPath   string `env:"PROJECT_ABS_PATH" env-required:"true"`
	ConfigPath       string `env:"CONFIG_PATH" env-required:"true"`
	PostgresUsername string `env:"POSTGRES_USERNAME" env-required:"true"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DBHost           string `env:"DB_HOST" env-required:"true"`
	PasswordHashSalt string `env:"PASSWORD_HASH_SALT" env-required:"true"`
	JWTSecret        string `env:"JWT_SECRET" env-required:"true"`

	ConfigAbsPath string
}

var (
	envCfgInst  = &EnvConfig{}
	loadEnvOnce = sync.Once{}
)

func Env() EnvConfig {
	loadEnvOnce.Do(func() {
		err := gconfig.ReadEnv(envFilename, envCfgInst)
		if err != nil {
			log.Fatalf("fatal reading env: %s\n", err)
		}

		envCfgInst.ConfigAbsPath = path.Join(envCfgInst.ProjectAbsPath, envCfgInst.ConfigPath)

		log.Infoln("Env successfully read")
	})

	return *envCfgInst
}
