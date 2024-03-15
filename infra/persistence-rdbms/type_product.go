package main

import "fmt"

const TableProducts = "products"

type Product struct {
	Code string `hera:"order:5"`

	CreatedAt int64 `hera:"columnname:created_at, order:2"`
	UpdatedAt int64 `hera:"columnname:updated_at, order:3"`
	DeletedAt int64 `hera:"columnname:deleted_at, order:4"`

	ID    uint `hera:"pk, order:1"`
	Price uint
}

func (p *Product) AsSQLInsert() string {
	return fmt.Sprintf(
		"insert into %s values ('%d', '%d', '%d', '%d', '%s', '%d');",
		TableProducts,
		p.ID,
		p.CreatedAt,
		p.UpdatedAt,
		p.DeletedAt,
		p.Code,
		p.Price,
	)
}

func (p *Product) AsSQLGetByPK(pk uint) string {
	return fmt.Sprintf(
		"select * from %s where id = %d",
		TableProducts,
		pk,
	)
}
