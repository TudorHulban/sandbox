package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	ddltable "github.com/TudorHulban/DDLTable"
	"github.com/stretchr/testify/require"
)

func TestStoredProcedures(t *testing.T) {
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

	fmt.Println(table.MigrationTable.Up)

	now := time.Now().UnixNano()

	testItem := Product{
		ID: 1,

		CreatedAt: int64(now),
		UpdatedAt: int64(now),

		Code:  "D43",
		Price: 100,
	}

	_, errCreateFN := pgxSimple.dbSimple.Exec(
		ctx,
		fnGetProductByPK,
	)
	require.NoError(t, errCreateFN)

	_, errInsert := pgxSimple.dbSimple.Exec(
		ctx,
		testItem.AsSQLInsert(),
	)
	require.NoError(t, errInsert)

	var reconstructedProduct Product

	errGetProduct := pgxSimple.GetProductByPKSimpleSP(
		ctx,
		1,
		&reconstructedProduct,
	)
	require.NoError(t, errGetProduct, errGetProduct)
	require.NotZero(t, reconstructedProduct)
	require.Equal(t, testItem.Price, reconstructedProduct.Price)
}
