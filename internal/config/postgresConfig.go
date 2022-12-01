package config

import (
	"path"
	"sync"
	"time"
)

const postgresConfigFilename = "postgres.config.yaml"

type PostgresConfig struct {
	MaxConnections          int    `yaml:"maxConnections" env-required:"true"`
	MaxConnectionAttempts   int    `yaml:"maxConnectionAttempts" env-required:"true"`
	WaitTimeoutMilliseconds int    `yaml:"waitTimeoutMilliseconds" env-required:"true"`
	Host                    string `yaml:"host" env-required:"true"`
	Port                    string `yaml:"port" env-required:"true"`
	Database                string `yaml:"database" env-required:"true"`

	Schema       string `yaml:"schema" env-default:"public"`
	UsersTable   string `yaml:"usersTable" env-default:"users"`
	TicketsTable string `yaml:"ticketsTable" env-default:"tickets"`
	FlightsTable string `yaml:"flightsTable" env-default:"flights"`

	Username    string
	Password    string
	WaitTimeout time.Duration
}

var (
	postgresCfgInst  = &PostgresConfig{}
	loadPostgresOnce = sync.Once{}
)

func Postgres() PostgresConfig {
	loadPostgresOnce.Do(func() {
		env := Env()
		readConfig(path.Join(env.ConfigAbsPath, postgresConfigFilename), postgresCfgInst)

		postgresCfgInst.Username = env.PostgresUsername
		postgresCfgInst.Password = env.PostgresPassword
		postgresCfgInst.WaitTimeout = time.Millisecond * time.Duration(postgresCfgInst.WaitTimeoutMilliseconds)
	})

	return *postgresCfgInst
}
