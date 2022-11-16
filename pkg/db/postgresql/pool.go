package postgresql

import (
	"api-app/pkg/utils"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
	Close()
}

func NewPool(ctx context.Context, cc *ClientConfig) (Pool, error) {
	config, err := NewConfig(cc)
	if err != nil {
		return nil, err
	}
	pool, err := utils.TryNTimesWaiting[Pool](cc.MaxAttempts, cc.WaitingDuration, func() (Pool, error) {
		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			return pool, err
		}

		err = pool.Ping(ctx)
		return pool, err
	})
	if err != nil {
		return nil, err
	}

	return pool, err
}
