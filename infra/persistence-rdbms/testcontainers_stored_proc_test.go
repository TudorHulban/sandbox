package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	ddltable "github.com/TudorHulban/DDLTable"
	"github.com/TudorHulban/GolangSandbox/helpers"
	"github.com/stretchr/testify/require"
)

func TestStoredProcedures(t *testing.T) {
	ctx := context.Background()

	configContainer, funcClean := GetTestContainerPG(
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

	table, errNew := ddltable.NewTable(
		_DDLRoot,
		&Product{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

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
				configContainer.AsDSNPGX(),
			)
		},
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

		CreatedAt: now,
		UpdatedAt: now,

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
