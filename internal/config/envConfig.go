package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"path"
	"sync"
)

// EnvConfig Env
type EnvConfig struct {
	ProjectAbsPath   string `env:"PROJECT_ABS_PATH" env-required:"true"`
	ConfigPath       string `env:"CONFIG_PATH" env-required:"true"`
	PostgresUsername string `env:"POSTGRES_USERNAME" env-required:"true"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" env-required:"true"`

	ConfigAbsPath string
}

var (
	envCfgInst  = &EnvConfig{}
	loadEnvOnce = sync.Once{}
)

func Env() EnvConfig {
	loadEnvOnce.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("fatal loading .env: %s\n", err.Error())
		}
		err = cleanenv.ReadEnv(envCfgInst)
		if err != nil {
			log.Fatalf("fatal reading env: %s\n", err.Error())
		}
		envCfgInst.ConfigAbsPath = path.Join(envCfgInst.ProjectAbsPath, envCfgInst.ConfigPath)
	})

	return *envCfgInst
}
