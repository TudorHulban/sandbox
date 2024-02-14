package main

import (
	"context"
	"testing"
	"time"

	testcontainerflyway "github.com/TudorHulban/GolangSandbox/infra/testcontainer-flyway"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// GetTestContainerPG returns test container and cleanup function for the created container.
func GetTestContainerPG(t *testing.T) (*postgres.PostgresContainer, func()) {
	ctx := context.Background()

	result, errContainer := postgres.RunContainer(
		ctx,

		testcontainers.WithImage("postgres:15.3-alpine"),

		postgres.WithDatabase(paramsPG.DBName),
		postgres.WithUsername(paramsPG.User),
		postgres.WithPassword(paramsPG.Password),

		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	require.NoError(t, errContainer)

	return result,
		func() {
			require.NoError(t,
				result.Terminate(ctx),
			)
		}
}

type ParamsGetTestContainerFlyway struct {
	MigrationsAbsolutePath string
	ExternalPGPort         uint
}

func GetTestContainerFlyway(params *ParamsGetTestContainerFlyway, t *testing.T) (*testcontainerflyway.ContainerFlyway, func()) {
	ctx := context.Background()

	result, errContainer := testcontainerflyway.RunContainer(
		ctx,

		&testcontainerflyway.ConfigFlyway{
			MigrationsPath: params.MigrationsAbsolutePath,

			DBName:   paramsPG.DBName,
			User:     paramsPG.User,
			Password: paramsPG.Password,

			Port: params.ExternalPGPort,
		},
		testcontainers.WithWaitStrategy(
			wait.ForLog("Flyway OSS Edition").
				WithStartupTimeout(10*time.Second),
		),
	)
	require.NoError(t, errContainer)

	return result,
		func() {
			require.NoError(t,
				result.Terminate(ctx),
			)
		}
}
