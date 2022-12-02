package postgresql

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type PGXPool interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
	Close()
}

func NewPGXPool(ctx context.Context, cc *ClientConfig) (PGXPool, error) {
	config, err := NewConfig(cc)

	if err != nil {
		return nil, err
	}

	pool, err := TryNTimesWaiting(ctx, cc.MaxConnectionAttempts, cc.WaitingDuration, func() (PGXPool, error) {
		//var pool PGXPool
		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			return nil, err
		}

		err = pool.Ping(ctx)
		if err != nil {
			return nil, err
		}

		return pool, nil
	})

	if err != nil {
		return nil, err
	}
	if pool == nil {
		return nil, errors.New("error couldn't connect to db")
	}
	return pool, err
}

func TryNTimesWaiting(ctx context.Context, n int, waitingDuration time.Duration, f func() (PGXPool, error)) (PGXPool, error) {
	var err error
	for i := 0; i < n; i++ {
		res, err := f()
		if err == nil {
			return res, nil
		}
		select {
		case <-time.After(waitingDuration):
			continue
		case <-ctx.Done():
			return nil, errors.New("error context was closed so")
		}
	}

	return nil, err
}
