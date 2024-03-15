package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/TudorHulban/GolangSandbox/helpers"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func migrationFiles(pathSource, pathTarget string, t *testing.T) []testcontainers.ContainerFile {
	entries, errRead := os.ReadDir(pathSource)
	require.NoError(t, errRead)

	containerFiles := make([]testcontainers.ContainerFile, len(entries), len(entries))

	for ix, entry := range entries {
		containerFiles[ix] = testcontainers.ContainerFile{
			HostFilePath:      pathSource + "/" + entry.Name(),
			ContainerFilePath: pathTarget + "/" + entry.Name(),
			FileMode:          0o777,
		}
	}

	return containerFiles
}

func TestFlyway(t *testing.T) {
	ctx := context.Background()

	cfg, funcClean := GetTestContainerPG(
		&ParamsGetTestContainerPostgres{
			PostgresCredentials: PostgresCredentials{
				DBPassword: paramsPG.DBPassword,
				DBUser:     paramsPG.DBUser,
				DBName:     paramsPG.DBName,
			},
		},
		t,
	)

	t.Cleanup(
		funcClean,
	)

	pgxSimple, errPGX := helpers.Retry[dbPGXSimple](
		&helpers.ParamsRetry{
			NoRetries: 5,
			FNRetry: func(numberRetry uint) time.Duration {
				return time.Millisecond * 50 * time.Duration(numberRetry+1)
			},
		},

		func() (*dbPGXSimple, error) {
			return NewPGXSimpleFromTestContainersDSN(
				ctx,
				cfg.AsDSNPGX(),
			)
		},
	)
	require.NoError(t, errPGX)
	require.NotNil(t, pgxSimple)

	RunTestContainerFlyway(
		&ParamsGetTestContainerFlyway{
			DBInfo:          cfg.DBInfo,
			MigrationsFiles: migrationFiles("../migrations", "/flyway/sql", t),
			ExternalPGPort:  uint(cfg.Port),
		},
		t,
	)

	count, errCount := pgxSimple.GetCountFromTable(ctx, "accounts")
	require.NoError(t, errCount)
	require.Zero(t, count)
}
