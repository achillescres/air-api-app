package config

import (
	"path"
	"sync"
)

const appConfigFile = "app.config.yaml"

// AppConfig App
type AppConfig struct {
	IsDev bool `yaml:"isDev" env-required:"true"`
	HTTP  struct {
		IP   string `yaml:"IP" env-default:"127.0.0.1"`
		Port string `yaml:"port" env-default:"7771"`
	} `yaml:"HTTP" env-required:"true"`
}

var (
	appCfgInst     = &AppConfig{}
	loadAppCfgOnce = sync.Once{}
)

func App() AppConfig {
	loadAppCfgOnce.Do(func() {
		appCfgPath := path.Join(Env().ConfigAbsPath, appConfigFile)
		readConfig(appCfgPath, appCfgInst)
	})

	return *appCfgInst
}
