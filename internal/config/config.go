package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Env struct {
	mainConfigPath string `env:"MAIN_CONFIG_PATH"`
}

type MainConfig struct {
	IsDev  *bool `yaml:"isDev"`
	Listen struct {
		Ip   string `yaml:"ip" env-default:"127.0.0.1"`
		Port string `yaml:"port" env-default:"7771"`
	} `yaml:"listen" env-required:"true"`
}

var (
	env      = &Env{}
	mainConf = &MainConfig{}
)

func init() {
	GetEnv()
	GetMainConfig()
}

var envOnce sync.Once
var mainConfigOnce sync.Once

func GetEnv() *Env {
	envOnce.Do(func() {
		log.Info("reading env")
		err := cleanenv.ReadEnv(env)
		if err != nil {
			log.Fatalf("fatal reading env: %s", err.Error())
		}
	})

	return env
}

func GetMainConfig() *MainConfig {
	mainConfigOnce.Do(func() {
		log.Info("reading config: MainConfig")
		localEnv := GetEnv()
		err := cleanenv.ReadConfig(localEnv.mainConfigPath, mainConf)
		if err != nil {
			log.Fatalf("fatal reading config: MainConfig: %s", err.Error())
		}
	})

	return mainConf
}
