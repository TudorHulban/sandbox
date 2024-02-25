package main

import (
	"context"
	"testing"

	ddltable "github.com/TudorHulban/DDLTable"
	"github.com/stretchr/testify/require"
)

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
		_DDLRoot,
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
		_DDLRoot,
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
