package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type dbPGXSimple struct {
	dbSimple *pgx.Conn
}

type dbPGX struct {
	dbSimple *pgx.Conn
	dbPool   *pgxpool.Pool
}

func NewPGXSimpleFromTestContainersDSN(ctx context.Context, cfg string) (*dbPGXSimple, error) {
	connPGXSimple, errConnPGXSimple := pgx.Connect(
		ctx,
		cfg,
	)
	if errConnPGXSimple != nil {
		return nil,
			errConnPGXSimple
	}

	return &dbPGXSimple{
			dbSimple: connPGXSimple,
		},
		nil
}

func NewDBPGX(ctx context.Context, cfg *ConfigPostgres) (*dbPGX, error) {
	connPGXPool, errConnPGXPool := pgxpool.New(
		ctx,
		cfg.AsDSNPGX(),
	)
	if errConnPGXPool != nil {
		return nil,
			errConnPGXPool
	}

	connPGXSimple, errConnPGXSimple := pgx.Connect(
		ctx,
		cfg.AsDSNPGX(),
	)
	if errConnPGXSimple != nil {
		return nil,
			errConnPGXSimple
	}

	return &dbPGX{
			dbSimple: connPGXSimple,
			dbPool:   connPGXPool,
		},
		nil
}

func (pgx *dbPGXSimple) GetProductByPKSimple(ctx context.Context, pk int, result *Product) error {
	return pgx.dbSimple.QueryRow(
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

func (pgx *dbPGX) GetProductByPKSimple(ctx context.Context, pk int, result *Product) error {
	return pgx.dbSimple.QueryRow(
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

func (pgx *dbPGXSimple) GetProductByPKSimpleSP(ctx context.Context, pk int, result *Product) error {
	return pgx.dbSimple.QueryRow(
		ctx,
		fmt.Sprintf(
			"select * from fnGetProduct(%d)",
			pk,
		),
	).Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.Code,
		&result.Price,
	)
}

func (pgx *dbPGX) GetProductByPKPool(ctx context.Context, pk int, result *Product) error {
	return pgx.dbPool.QueryRow(
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
