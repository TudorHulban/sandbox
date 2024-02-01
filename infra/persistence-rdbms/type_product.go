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
function fnGetProduct(in pk int)
returns table
            (
            	id bigint,
            	created_at bigint,
                updated_at bigint,
				deleted_at bigint,
				code text,
				price bigint
            )
language plpgsql
as 
$$
begin
	return query
	select
	p.id,
	p.created_at,
	p.updated_at,
	p.deleted_at,
	p.code,
	p.price
from
	products p
where
	p.id = pk;
end
$$;
`

	return fn
}
