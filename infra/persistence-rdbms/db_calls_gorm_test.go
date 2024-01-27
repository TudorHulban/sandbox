package main

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		Code:  "D42",
		Price: 100,
	}

	db.Create(
		&testItem,
	)

	connSQLite := NewDBGORM(db)

	var reconstructed Product

	connSQLite.FirstByPK(
		1,
		&reconstructed,
	)

	require.NotZero(t, reconstructed)
	require.Equal(t,
		testItem.Price,
		reconstructed.Price,
	)
}
