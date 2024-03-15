package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/TudorHulban/GolangSandbox/helpers"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _dbConnection *pgxpool.Pool
var mu sync.RWMutex

func GetPostgresDBConnection(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	mu.RLock()
	if _dbConnection != nil {
		defer mu.RUnlock()

		return _dbConnection, nil
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()

	_dbConnection, errConnect := helpers.Retry[pgxpool.Pool](
		&helpers.ParamsRetry{
			NoRetries: 5,
			FNRetry: func(numberRetry uint) time.Duration {
				return time.Millisecond * 50 * time.Duration(numberRetry+1)
			},
		},

		func() (*pgxpool.Pool, error) {
			return pgxpool.New(
				ctx,
				dsn,
			)
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
