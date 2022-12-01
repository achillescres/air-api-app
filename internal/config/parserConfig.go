package config

import (
	"path"
	"sync"
)

const taisParserConfigFilename = "parser.config.yaml"

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
		readConfig(path.Join(Env().ConfigAbsPath, taisParserConfigFilename), taisParserCfgInst)
	})

	return *taisParserCfgInst
}
