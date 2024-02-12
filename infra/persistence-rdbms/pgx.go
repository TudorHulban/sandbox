package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/TudorHulban/GolangSandbox/apperrors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type retryInterval func(numberRetry uint) time.Duration

type paramsConnectPGXWithRetries struct {
	DSN       string
	NoRetries uint
	FNRetry   retryInterval
}

var _dbConnection *pgxpool.Pool
var mu sync.RWMutex

func connectPGXWithRetries(ctx context.Context, params *paramsConnectPGXWithRetries) (*pgxpool.Pool, error) {
	var errResult error

	for i := 0; i < int(params.NoRetries); i++ {
		connPGXPool, errConnPGXPool := pgxpool.New(
			ctx,
			params.DSN,
		)
		if errConnPGXPool == nil {
			return connPGXPool, nil
		}

		if i == int(params.NoRetries) {
			errResult = apperrors.ErrCausedByInfrastructure{
				Caller:  "connectPGXWithRetries",
				Calling: "pgxpool.New",
				Issue: fmt.Errorf(
					"database not ready after %d retries:%w",
					params.NoRetries,
					errConnPGXPool,
				),
			}

			break
		}

		time.Sleep(
			params.FNRetry(uint(i)),
		)
	}

	return nil,
		errResult
}

func GetPostgresDBConnection(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	mu.RLock()
	if _dbConnection != nil {
		defer mu.RUnlock()

		return _dbConnection, nil
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()

	_dbConnection, errConnect := connectPGXWithRetries(
		ctx,
		&paramsConnectPGXWithRetries{
			DSN:       dsn,
			NoRetries: 3,
			FNRetry: func(numberRetry uint) time.Duration {
				return time.Millisecond * 1500 * time.Duration(numberRetry+1)
			},
		},
	)
	if errConnect != nil {
		return nil,
			fmt.Errorf(
				"GetPostgresDBConnection connectDBPostgres:%w",
				errConnect,
			)
	}

	return _dbConnection, nil
}
