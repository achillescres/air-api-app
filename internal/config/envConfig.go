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
	ProjectPath   string `env:"PROJECTPATH" env-required:"true"`
	ConfigPath    string `env:"CONFIGPATH" env-required:"true"`
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
			log.Fatalf("fatal loading .env: %s", err.Error())
		}
		err = cleanenv.ReadEnv(envCfgInst)
		if err != nil {
			log.Fatalf("fatal reading env: %s", err.Error())
		}
		envCfgInst.ConfigAbsPath = path.Join(envCfgInst.ProjectPath, envCfgInst.ConfigPath)
	})

	return *envCfgInst
}
