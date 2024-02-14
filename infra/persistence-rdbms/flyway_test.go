package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/require"
)

func TestFlyway(t *testing.T) {
	ctx := context.Background()

	pgContainer, funcCleanPg := GetTestContainerPG(t)

	t.Cleanup(
		funcCleanPg,
	)

	connDSN, errConnection := pgContainer.ConnectionString(
		ctx,
		"sslmode=disable",
	)
	require.NoError(t, errConnection)

	pgxSimple, errPGX := NewPGXSimpleFromTestContainersDSN(
		ctx,
		connDSN,
	)
	require.NoError(t, errPGX)
	require.NotNil(t, pgxSimple)

	portExternalPGContainer, errPort := pgContainer.MappedPort(
		ctx,
		nat.Port("5432"),
	)
	require.NoError(t, errPort)
	require.NotZero(t, portExternalPGContainer)

	fmt.Println("container port:", portExternalPGContainer)

	flywayContainer, funcCleanFlyway := GetTestContainerFlyway(
		&ParamsGetTestContainerFlyway{
			MigrationsAbsolutePath: "../migrations",
			ExternalPGPort:         uint(portExternalPGContainer.Int()),
		},
		t,
	)

	t.Cleanup(
		funcCleanFlyway,
	)

	require.NotNil(t, flywayContainer)

	count, errCount := pgxSimple.GetCountFromTable(ctx, "accounts")
	require.NoError(t, errCount)
	require.Zero(t, count)
}
