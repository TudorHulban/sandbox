package main

const fnInsertProduct = `
create or replace
function fnInsertProduct(
in created_at bigint,
in updated_at bigint,

in code text,
in price bigint)

returns bigint
language plpgsql
as 
$$
declare 
inserted_id bigint;

begin
insert
	into
	products(
	id,
	created_at,
	updated_at,
	code,
	price)
values (
nextval('products_id_seq'),
created_at,
updated_at,
code,
price)
returning id
into
	inserted_id;

return inserted_id;
end
$$;
`

const fnGetProductByPK = `
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
