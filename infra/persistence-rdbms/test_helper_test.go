package main

import (
	"context"
	"testing"
	"time"

	"github.com/TudorHulban/GolangSandbox/config"
	testcontainerflyway "github.com/TudorHulban/GolangSandbox/infra/testcontainer-flyway"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresCredentials struct {
	DBPassword string
	DBUser     string
	DBName     string
}

type ParamsGetTestContainerPostgres struct {
	PostgresCredentials
	ExternalPGPort string
}

func GetTestContainerPG(params *ParamsGetTestContainerPostgres, t *testing.T) (*config.ConfigPostgres, func()) {
	ctx := context.Background()

	reqContainer := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:15.3-alpine",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_PASSWORD": paramsPG.DBPassword,
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

	return &config.ConfigPostgres{
			DBInfo: config.DBInfo{
				DBName:     params.DBName,
				DBUser:     params.DBUser,
				DBPassword: params.DBPassword,
				Host:       hostPostgres,
				Port:       uint(portPostgres.Int()),
			},
		},

		func() {
			require.NoError(t,
				container.Terminate(ctx),
			)
		}
}

type ParamsGetTestContainerFlyway struct {
	config.DBInfo

	MigrationsFiles []testcontainers.ContainerFile
	ExternalPGPort  uint
}

func RunTestContainerFlyway(params *ParamsGetTestContainerFlyway, t *testing.T) {
	ctx := context.Background()

	errContainer := testcontainerflyway.RunContainerFlyway(
		ctx,

		&testcontainerflyway.ConfigFlyway{
			MigrationsFiles: params.MigrationsFiles,

			DBInfo: config.DBInfo{
				DBName:     params.DBName,
				DBUser:     params.DBUser,
				DBPassword: params.DBPassword,
				Host:       params.Host,
				Port:       params.ExternalPGPort,
			},
		},
		testcontainers.WithWaitStrategy(
			wait.ForLog("Flyway OSS Edition").
				WithStartupTimeout(10*time.Second),
		),
	)
	require.NoError(t, errContainer)
}
