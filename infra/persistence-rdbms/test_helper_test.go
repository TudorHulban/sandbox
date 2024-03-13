package main

import (
	"context"
	"testing"
	"time"

	testcontainerflyway "github.com/TudorHulban/GolangSandbox/infra/testcontainer-flyway"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type ParamsGetTestContainerPostgres struct {
	DBPassword string
	DBUser     string
	DBName     string

	ExternalPGPort string
}

// GetTestContainerPG returns test container and cleanup function for the created container.
func GetTestContainerPG(params *ParamsGetTestContainerPostgres, t *testing.T) (*testcontainers.Container, *ConfigPostgres, func()) {
	ctx := context.Background()

	reqContainer := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:15.3-alpine",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_PASSWORD": paramsPG.Password,
				"POSTGRES_USER":     params.DBUser,
				"POSTGRES_DB":       params.DBName,
			},
			NetworkMode: "host",
			WaitingFor:  wait.ForLog("database system is ready to accept connections"),
		},
		Started: true,
	}

	container, errCrContainer := testcontainers.GenericContainer(ctx, reqContainer)
	require.NoError(t, errCrContainer)

	portPostgres, errPort := container.MappedPort(ctx, nat.Port("5432"))
	require.NoError(t, errPort)

	hostPostgres, errHost := container.Host(ctx)
	require.NoError(t, errHost)

	return &container,
		&ConfigPostgres{
			DBName:   params.DBName,
			User:     params.DBUser,
			Password: params.DBPassword,
			Host:     hostPostgres,
			Port:     portPostgres.Int(),
		},
		func() {
			require.NoError(t,
				container.Terminate(ctx),
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
