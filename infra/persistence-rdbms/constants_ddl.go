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

// ex. call spupdateproduct(2,null,null,0,'new text', 399)
const spUpdateProduct = `
create or replace
procedure spupdateproduct(
in pk bigint,
in pCreatedAt bigint,
in pUpdatedAt bigint,
in pDeletedAt bigint,
in pCode text,
in pPrice bigint
)

language plpgsql
as 
$$
begin
update 
	products p
set 
	created_at = coalesce (pCreatedAt, p.created_at),
	updated_at = coalesce (pUpdatedAt, p.updated_at),
	deleted_at = coalesce (pDeletedAt, p.deleted_at),
	code = coalesce (pCode, p.code),
	price = coalesce (pPrice, p.price)
where
	id = pk;
end
$$;
`

const spDeleteSoftProduct = `
create or replace
procedure spsoftdeleteproduct(
in pk bigint,
in deletedTime bigint
)

language plpgsql
as 
$$
begin
update
	products
set
	deleted_at = deletedTime
where
	id = pk;
end
$$;
`

const spDeleteHardProduct = `
create or replace
procedure spharddeleteproduct(
in pk bigint
)

language plpgsql
as 
$$
begin
delete from 
	products
where
	id = pk;
end
$$;
`
