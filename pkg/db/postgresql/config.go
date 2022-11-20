package postgresql

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type ClientConfig struct {
	MaxConnections                           int
	MaxConnectionAttempts                    int
	WaitingDuration                          time.Duration
	Username, Password, Host, Port, Database string
}

func NewConfig(cc *ClientConfig) (*pgxpool.Config, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cc.Username, cc.Password, cc.Host, cc.Port, cc.Database)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	addOptionsToConfig(cc, config)

	return config, err
}

func addOptionsToConfig(cc *ClientConfig, config *pgxpool.Config) *pgxpool.Config {
	config.MaxConns = int32(cc.MaxConnections)
	return config
}
