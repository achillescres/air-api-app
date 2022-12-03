package postgresql

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
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

	pool, err := TryNTimesWaiting(ctx, cc.MaxConnectionAttempts, cc.WaitingDuration, func(ctx context.Context) (PGXPool, error) {
		log.Infof("trying to connect to db: %s\n", config.ConnString())
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

func TryNTimesWaiting(ctx context.Context, n int, waitingDuration time.Duration, f func(ctx context.Context) (PGXPool, error)) (PGXPool, error) {
	var res PGXPool
	var err error

	var ctxTimeout context.Context
	var cancel context.CancelFunc
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	for i := 0; i < n; i++ {
		ctxTimeout, cancel = context.WithTimeout(context.Background(), waitingDuration)

		go func() {
			res, err = f(ctxTimeout)
		}()
		select {
		case <-ctxTimeout.Done():
			if res == nil || err != nil {
				continue
			}
			cancel()
			return res, nil
		case <-ctx.Done():
			cancel()
			return nil, errors.New("error global context was closed so")
		}
	}

	if cancel != nil {
		cancel()
	}
	return res, err
}
