package config

import (
	"path"
	"sync"
)

const taisConfigFilename = "tais.config.yaml"

type TaisConfig struct {
	TaisDirPath                 string `yaml:"taisDirPath" env-required:"true"`
	TotalCashDelimiterIndex     int    `yaml:"totalCashDelimiterIndex" env-required:"true"`
	DefaultFlightsSliceCapacity int    `yaml:"defaultFlightsSliceCapacity" env-required:"true"`

	TaisAbsDirPath string
}

var (
	taisCfgInst     = &TaisConfig{}
	loadTaisCfgOnce = sync.Once{}
)

func Tais() TaisConfig {
	loadTaisCfgOnce.Do(func() {
		readConfig(path.Join(Env().ConfigAbsPath, taisConfigFilename), taisCfgInst)

		taisCfgInst.TaisAbsDirPath = path.Join(Env().ProjectAbsPath, taisCfgInst.TaisDirPath)
	})

	return *taisCfgInst
}
