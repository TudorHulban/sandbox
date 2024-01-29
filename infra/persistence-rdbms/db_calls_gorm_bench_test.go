package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func BenchmarkGORMSQLite(b *testing.B) {
	db, errOpen := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{},
	)
	require.NoError(b, errOpen)
	require.NotNil(b, db)

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

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		connSQLite.FirstByPK(
			1,
			&reconstructed,
		)
	}
}

func BenchmarkGORMRAWPG(b *testing.B) {
	connGORM, errConnGORM := gorm.Open(
		postgres.Open(
			paramsPG.AsDSNGORM(),
		),
		&gorm.Config{
			DisableAutomaticPing: true,
		},
	)
	require.NoError(b, errConnGORM)

	gormPG := NewDBGORM(connGORM)

	var reconstructed Product

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		gormPG.GetProductByPK(
			1,
			&reconstructed,
		)
	}
}

func BenchmarkGORMPG(b *testing.B) {
	connGORM, errConnGORM := gorm.Open(
		postgres.Open(
			paramsPG.AsDSNGORM(),
		),
		&gorm.Config{
			DisableAutomaticPing: true,
		},
	)
	require.NoError(b, errConnGORM)

	gormPG := NewDBGORM(connGORM)

	var reconstructed Product

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		gormPG.FirstByPK(
			1,
			&reconstructed,
		)
	}
}
