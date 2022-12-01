package config

import (
	"path"
	"sync"
)

const handlerConfigFilename = "usecase.config.yaml"

type HandlerConfig struct {
	DefaultTableCapacity int `yaml:"defaultTableCapacity" env-default:"16"`
}

var (
	handlerCfgInst     = &HandlerConfig{}
	loadHandlerCfgOnce = sync.Once{}
)

func Handler() HandlerConfig {
	loadHandlerCfgOnce.Do(func() {
		readConfig(path.Join(Env().ConfigAbsPath, handlerConfigFilename), handlerCfgInst)
	})

	return *handlerCfgInst
}
