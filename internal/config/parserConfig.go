package config

import (
	"path"
	"sync"
)

const taisParserConfigFile = "parser.config.yaml"

type TaisParserConfig struct {
	TaisDirPath             string `yaml:"taisDirPath" env-required:"true"`
	TotalCashDelimiterIndex int    `yaml:"totalCashDelimiterIndex" env-required:"true"`
}

var (
	taisParserCfgInst     = &TaisParserConfig{}
	loadTaisParserCfgOnce = sync.Once{}
)

func TaisParser() TaisParserConfig {
	loadTaisParserCfgOnce.Do(func() {
		taisParserCfgPath := path.Join(Env().ConfigAbsPath, taisParserConfigFile)
		readConfig(taisParserCfgPath, taisParserCfgInst)
	})

	return *taisParserCfgInst
}
