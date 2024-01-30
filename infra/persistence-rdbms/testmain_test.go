package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	ddltable "github.com/TudorHulban/GolangSandbox/infra/ddl-table"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestContainer1(t *testing.T) {
	ctx := context.Background()

	pgContainer, errContainer := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),

		postgres.WithDatabase(paramsPG.DBName),
		postgres.WithUsername(paramsPG.User),
		postgres.WithPassword(paramsPG.Password),

		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	require.NoError(t, errContainer)

	t.Cleanup(
		func() {
			fmt.Println("cleanup test....")

			require.NoError(t,
				pgContainer.Terminate(ctx),
			)
		},
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
		table.AsDDLPostgres(),
	)
	require.NoError(t, errExec)

	testItem := Product{
		ID:    1,
		Code:  "D43",
		Price: 100,
	}

	tag, errInsert := pgxSimple.dbSimple.Exec(
		ctx,
		testItem.AsSQLInsert(),
	)
	require.NoError(t, errInsert)

	fmt.Println("tag:", tag, testItem.AsSQLInsert())

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
