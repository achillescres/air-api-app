package config

import (
	"path"
	"sync"
	"time"
)

const dbConfigFile = "db.config.yaml"

type DBConfig struct {
	MaxConnectionAttempts   int    `yaml:"maxConnectionAttempts"`
	WaitTimeoutMilliseconds int    `yaml:"waitTimeoutMilliseconds"`
	Username                string `yaml:"username"`
	Password                string `yaml:"password"`
	Host                    string `yaml:"host"`
	Port                    string `yaml:"port"`
	Database                string `yaml:"database"`
	WaitTimeout             time.Duration

	Schema       string `yaml:"schema" env-default:"public"`
	UsersTable   string `yaml:"usersTable" env-default:"users"`
	TicketsTable string `yaml:"ticketsTable" env-default:"tickets"`
	FloghtsTable string `yaml:"flightsTable" env-default:"flights"`
}

var (
	dbCfgInst  = &DBConfig{}
	loadDBOnce = sync.Once{}
)

func DB() DBConfig {
	loadDBOnce.Do(func() {
		dbCfgPath := path.Join(Env().ConfigAbsPath, dbConfigFile)
		readConfig(dbCfgPath, dbCfgInst)
		dbCfgInst.WaitTimeout = time.Millisecond * time.Duration(dbCfgInst.WaitTimeoutMilliseconds)
	})

	return *dbCfgInst
}
