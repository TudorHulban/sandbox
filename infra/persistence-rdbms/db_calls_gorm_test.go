package main

import (
	"testing"
	"time"

	ddltable "github.com/TudorHulban/DDLTable"
	"github.com/TudorHulban/GolangSandbox/helpers"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGORM(t *testing.T) {
	db, errOpen := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{},
	)
	require.NoError(t, errOpen)
	require.NotNil(t, db)

	db.AutoMigrate(
		&Product{},
	)

	testItem := Product{
		Code:  "D43",
		Price: 100,
	}

	db.Create(
		&testItem,
	)

	connSQLite := NewDBGORM(db)

	var reconstructed1 Product

	connSQLite.FirstByPK(
		1,
		&reconstructed1,
	)

	require.NotZero(t, reconstructed1)
	require.Equal(t,
		testItem.Price,
		reconstructed1.Price,
	)

	var reconstructed2 Product

	connSQLite.GetProductByPK(
		1,
		&reconstructed2,
	)

	require.NotZero(t, reconstructed2)
	require.Equal(t,
		testItem.Price,
		reconstructed2.Price,
	)
}

func TestGORMDDL(t *testing.T) {
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

	connGORM, errConnGORM := helpers.Retry[gorm.DB](
		&helpers.ParamsRetry{
			NoRetries: 5,
			FNRetry: func(numberRetry uint) time.Duration {
				return time.Millisecond * 50 * time.Duration(numberRetry+1)
			},
		},

		func() (*gorm.DB, error) {
			return gorm.Open(
				postgres.Open(
					cfg.AsDSNGORM(),
				),
			)
		},
	)

	require.NoError(t, errConnGORM)

	table, errNew := ddltable.NewTable(
		_DDLRoot,
		&Product{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	require.NoError(t,
		connGORM.Exec(
			table.MigrationTable.Up,
		).
			Error,
	)

	testItem := Product{
		Code:  "D43",
		Price: 100,
	}

	require.NoError(t,
		connGORM.Create(
			&testItem,
		).
			Error,
	)

	var reconstructedProduct Product

	require.NoError(t,
		connGORM.First(
			&reconstructedProduct,
			1,
		).
			Error,
	)
	require.NotZero(t, reconstructedProduct)
	require.Equal(t,
		testItem.Price,
		reconstructedProduct.Price,
	)
}
