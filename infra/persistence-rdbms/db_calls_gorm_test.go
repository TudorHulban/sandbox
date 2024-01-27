package main

import (
	"testing"

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

// os: mint 21.3
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkGORM-16    	   21916	     54327 ns/op	    5555 B/op	     130 allocs/op
func BenchmarkGORM(b *testing.B) {
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

// os: mint 21.3
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkGORMPG-16    	   16459	     93467 ns/op	    5586 B/op	      98 allocs/op
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
