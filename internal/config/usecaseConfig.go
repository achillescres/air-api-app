package config

import (
	"path"
	"sync"
)

const usecaseConfigFile = "usecase.config.yaml"

// UsecaseConfig Usecase (layer)
type UsecaseConfig struct {
	DefaultTableCapacity int `yaml:"defaultTableCapacity" env-default:"16"`
}

var (
	usecaseCfgInst     = &UsecaseConfig{}
	loadUsecaseCfgOnce = sync.Once{}
)

func Usecase() UsecaseConfig {
	loadUsecaseCfgOnce.Do(func() {
		usecaseCfgPath := path.Join(Env().ConfigAbsPath, usecaseConfigFile)
		readConfig(usecaseCfgPath, usecaseCfgInst)
	})

	return *usecaseCfgInst
}
