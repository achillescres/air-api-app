package config

import (
	"path"
	"sync"
)

const tablesConfigFilename = "tables.config.yaml"

type TablesConfig struct {
	FlightTableDefaultCapacity int `yaml:"flightTableDefaultCapacity" env-default:"1024"`
}

var (
	tablesConfigInst     = &TablesConfig{}
	loadTablesConfigOnce = &sync.Once{}
)

func Tables() TablesConfig {
	loadTablesConfigOnce.Do(func() {
		readConfig(path.Join(Env().ConfigAbsPath, tablesConfigFilename), tablesConfigInst)
	})

	return *tablesConfigInst
}
