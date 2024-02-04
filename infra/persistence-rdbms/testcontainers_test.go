package main

import (
	"context"
	"testing"
	"time"

	ddltable "github.com/TudorHulban/GolangSandbox/infra/ddl-table"
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

func TestContainer1(t *testing.T) {
	ctx := context.Background()

	pgContainer, funcClean := GetTestContainerPG(t)

	t.Cleanup(
		funcClean,
	)

	connDSN, errConnection := pgContainer.ConnectionString(
		ctx,
		"sslmode=disable",
	)
	require.NoError(t, errConnection)

	table, errNew := ddltable.NewTable(
		&Product{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	pgxSimple, errPGX := NewPGXSimpleFromTestContainersDSN(
		ctx,
		connDSN,
	)
	require.NoError(t, errPGX)
	require.NotNil(t, pgxSimple)

	_, errCreateTable := pgxSimple.dbSimple.Exec(
		ctx,
		table.MigrationTable.Up,
	)
	require.NoError(t, errCreateTable)

	testItem := Product{
		ID:    1,
		Code:  "D43",
		Price: 100,
	}

	_, errInsert := pgxSimple.dbSimple.Exec(
		ctx,
		testItem.AsSQLInsert(),
	)
	require.NoError(t, errInsert)

	var reconstructedProduct Product

	require.NoError(t,
		pgxSimple.GetProductByPKSimple(
			ctx,
			1,
			&reconstructedProduct,
		),
	)
	require.NotZero(t, reconstructedProduct)
	require.Equal(t, testItem.Price, reconstructedProduct.Price)
}

func TestContainer2(t *testing.T) {
	ctx := context.Background()

	pgContainer, funcClean := GetTestContainerPG(t)

	t.Cleanup(
		funcClean,
	)

	connDSN, errConnection := pgContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, errConnection)

	table, errNew := ddltable.NewTable(
		&Product{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	pgxSimple, errPGX := NewPGXSimpleFromTestContainersDSN(ctx, connDSN)
	require.NoError(t, errPGX)
	require.NotNil(t, pgxSimple)

	_, errExec := pgxSimple.dbSimple.Exec(
		ctx,
		table.MigrationTable.Up,
	)
	require.NoError(t, errExec)

	testItem := Product{
		ID:    1,
		Code:  "D43",
		Price: 100,
	}

	_, errInsert := pgxSimple.dbSimple.Exec(
		ctx,
		testItem.AsSQLInsert(),
	)
	require.NoError(t, errInsert)

	var reconstructedProduct Product

	require.NoError(t,
		pgxSimple.GetProductByPKSimple(
			ctx,
			1,
			&reconstructedProduct,
		),
	)
	require.NotZero(t, reconstructedProduct)
	require.Equal(t, testItem.Price, reconstructedProduct.Price)
}
