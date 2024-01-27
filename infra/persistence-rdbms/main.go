package main

import (
	"context"
	"os"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	connPGX, errConnPGX := pgxpool.New(
		ctx,
		paramsPG.AsDSNPGX(),
	)
	if errConnPGX != nil {
		os.Exit(10)
	}

	driverMigrations, errDriver := postgres.WithInstance(connPGX, &postgres.Config{})
}
