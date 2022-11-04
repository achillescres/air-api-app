package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"reflect"
	"sync"
)

type Config interface{}

// EnvConfig Env
type EnvConfig struct {
	AppConfigPath     string `env:"mainConfigPath"`
	UsecaseConfigPath string `env:"usecaseConfigPath"`
	ParserConfigPath  string `env:"parserConfigPath"`
}

var _ Config = (*EnvConfig)(nil)

type EnvConfigInvoker func() EnvConfig

// AppConfig App
type AppConfig struct {
	IsDev  bool `yaml:"isDev"`
	Listen struct {
		IP   string `yaml:"ip" env-default:"127.0.0.1"`
		Port string `yaml:"port" env-default:"7771"`
	} `yaml:"listen" env-required:"true"`
}

var _ Config = (*AppConfig)(nil)

type AppConfigInvoker func() AppConfig

// UsecaseConfig Usecase (layer)
type UsecaseConfig struct {
	TableCapacity int `yaml:"defaultCapacity" env-default:"16"`
}

var _ Config = (*UsecaseConfig)(nil)

type UsecaseConfigInvoker func() UsecaseConfig

// ParserConfig Parser
type ParserConfig struct {
	TaisFilePath string `yaml:"taisFilePath"`
}

var _ Config = (*ParserConfig)(nil)

type ParserConfigInvoke func() ParserConfig

func init() {
	loadAllCfgs()
}

var (
	envCfgInst     *EnvConfig
	appCfgInst     *AppConfig
	usecaseCfgInst *UsecaseConfig
	parserCfgInst  *ParserConfig
)

var (
	loadOnce sync.Once
	loaded   = false
)

func readConfig(path string, inst Config) {
	cfgName := reflect.TypeOf(path).String()
	log.Infof("reading config: %s\n", cfgName)
	if err := cleanenv.ReadConfig(path, inst); err != nil {
		log.Fatalf("fatal reading config %s: %s\n", cfgName, err.Error())
	}
}

func loadAllCfgs() {
	loadOnce.Do(func() {
		log.Info("reading config EnvConfig(.env)")
		err := cleanenv.ReadEnv(envCfgInst)
		if err != nil {
			log.Fatalf("fatal reading env: %s\n", err.Error())
		}

		readConfig(envCfgInst.AppConfigPath, appCfgInst)
		readConfig(envCfgInst.UsecaseConfigPath, usecaseCfgInst)
		readConfig(envCfgInst.ParserConfigPath, parserCfgInst)

		loaded = true
	})
}

var fabricsBorn = false

func OnceGetConfigInvokes() (EnvConfigInvoker, AppConfigInvoker, UsecaseConfigInvoker, ParserConfigInvoke, error) {
	loadAllCfgs()
	if fabricsBorn {
		log.Fatalln("fatal config fabrics are already born")
	}
	if !loaded {
		return nil, nil, nil, nil, errors.New("error configs aren't loaded")
	}

	fabricsBorn = true

	invokeEnvCfg := func() EnvConfig {
		return *envCfgInst
	}

	invokeAppCfg := func() AppConfig {
		return *appCfgInst
	}

	invokeUsecaseCfg := func() UsecaseConfig {
		return *usecaseCfgInst
	}

	invokeParserCfg := func() ParserConfig {
		return *parserCfgInst
	}

	return invokeEnvCfg, invokeAppCfg, invokeUsecaseCfg, invokeParserCfg, nil
}
