package config

import (
	"path"
	"sync"
)

const handlerConfigFile = "usecase.config.yaml"

type HandlerConfig struct {
	DefaultTableCapacity int `yaml:"defaultTableCapacity" env-default:"16"`
}

var (
	handlerCfgInst     = &HandlerConfig{}
	loadHandlerCfgOnce = sync.Once{}
)

func Handler() HandlerConfig {
	loadHandlerCfgOnce.Do(func() {
		handlerCfgPath := path.Join(Env().ConfigAbsPath, handlerConfigFile)
		readConfig(handlerCfgPath, handlerCfgInst)
	})

	return *handlerCfgInst
}
