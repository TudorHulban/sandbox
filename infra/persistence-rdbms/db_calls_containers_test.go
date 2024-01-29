package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	ddltable "github.com/TudorHulban/GolangSandbox/infra/ddl-table"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func TestContainersPG(t *testing.T) {
	pool, errDockerTest := dockertest.NewPool("")
	require.NoError(t, errDockerTest)

	errPing := pool.Client.Ping()
	require.NoError(t, errPing)

	resource, errContainer := pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        "11",
			Env: []string{
				fmt.Sprintf("POSTGRES_USER=%s", paramsPG.User),
				fmt.Sprintf("POSTGRES_PASSWORD=%s", paramsPG.Password),
				fmt.Sprintf("POSTGRES_DB=%s", paramsPG.DBName),
				fmt.Sprintf("listen_addresses=%d", paramsPG.Port),
			},
		},

		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		},
	)
	require.NoError(t, errContainer)

	defer resource.Expire(60)

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		paramsPG.User,
		paramsPG.Password,
		paramsPG.DBName,
		hostAndPort,
	)

	pool.MaxWait = 20 * time.Second

	errRetry := pool.Retry(
		func() error {
			db, errOpen := sql.Open("postgres", databaseUrl)
			require.NoError(t, errOpen)

			return db.Ping()
		},
	)
	require.NoError(t, errRetry)

	connGORM, errConnGORM := gorm.Open(
		postgres.Open(
			paramsPG.AsDSNGORM(),
		),
		&gorm.Config{
			DisableAutomaticPing: true,
		},
	)
	require.NoError(t, errConnGORM)

	table, errNew := ddltable.NewTable(
		&Product{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	require.NoError(t,
		connGORM.Exec(
			table.AsDDLPostgres(),
		).
			Error,
	)

	testItem := Product{
		Code:  "D43",
		Price: 100,
	}

	require.NoError(t,
		connGORM.Create(
			&testItem,
		).
			Error,
	)
}
