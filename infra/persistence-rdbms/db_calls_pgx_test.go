package main

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

// os: mint 21.3
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkPGX-16    	   33958	     63741 ns/op	     560 B/op	      12 allocs/op
func BenchmarkPGX(b *testing.B) {
	ctx := context.Background()

	connPGX, errConnPGX := pgx.Connect(
		ctx,
		paramsPG.AsDSNPGX(),
	)
	require.NoError(b, errConnPGX)

	pgxPG := NewDBPGX(connPGX)

	var reconstructed Product

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		pgxPG.GetProductByPK(
			ctx,
			1,
			&reconstructed,
		)
	}
}
