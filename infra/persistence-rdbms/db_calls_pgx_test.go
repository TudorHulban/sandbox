package main

import (
	"context"
	"testing"

	"github.com/TudorHulban/GolangSandbox/config"
	"github.com/stretchr/testify/require"
)

func BenchmarkPGXSimple(b *testing.B) {
	ctx := context.Background()

	connPGX, errConnPGX := NewDBPGX(
		ctx,
		&config.ConfigPostgres{
			DBInfo: paramsPG,
		},
	)
	require.NoError(b, errConnPGX)

	var reconstructed Product

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		connPGX.GetProductByPKSimple(
			ctx,
			1,
			&reconstructed,
		)
	}
}

func BenchmarkPGXPool(b *testing.B) {
	ctx := context.Background()

	connPGX, errConnPGX := NewDBPGX(
		ctx,
		&config.ConfigPostgres{
			DBInfo: paramsPG,
		},
	)
	require.NoError(b, errConnPGX)

	var reconstructed Product

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		connPGX.GetProductByPKPool(
			ctx,
			1,
			&reconstructed,
		)
	}
}
