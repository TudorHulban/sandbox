package main

import (
	"testing"

	ddltable "github.com/TudorHulban/GolangSandbox/infra/ddl-table"
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
	connGORM, errConnGORM := gorm.Open(
		postgres.Open(
			paramsPG.AsDSNGORM(),
		),
		&gorm.Config{
			DisableAutomaticPing: true,
		},
	)
	require.NoError(t, errConnGORM)

	table, errNew := ddltable.NewTable(
		&Product{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	require.NoError(t,
		connGORM.Exec(
			table.AsDDLPostgres(),
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
}
