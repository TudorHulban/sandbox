package main

import "fmt"

const TableProducts = "products"

type Product struct {
	// gorm.Model
	ID        uint  `hera:"pk"`
	CreatedAt int64 `hera:"columnname:created_at"`
	UpdatedAt int64 `hera:"columnname:updated_at"`
	DeletedAt int64 `hera:"columnname:deleted_at"`

	Code  string
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

func (p *Product) DDLFunctionGetByPK() string {
	fn := `
	create or replace
	function fnGetProduct(
	in pk int, 
	
	out id int,
	out created_at int,
	out updated_at int)
	as $$
	select
		id,
		created_at,
		updated_at
	from
		products
	where
		id = $1 $$
	language sql;
`

	return fn
}
