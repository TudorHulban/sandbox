package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type dbPGX struct {
	db *pgx.Conn
}

func NewDBPGX(db *pgx.Conn) *dbPGX {
	return &dbPGX{
		db: db,
	}
}

func (pgx *dbPGX) GetProductByPK(ctx context.Context, pk int, result *Product) error {
	return pgx.db.QueryRow(
		ctx,
		"select id, created_at, updated_at, deleted_at, code, price from products where id=$1",
		pk,
	).
		Scan(
			&result.ID,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt,
			&result.Code,
			&result.Price,
		)
}
