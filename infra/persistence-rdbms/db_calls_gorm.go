package main

import "gorm.io/gorm"

type dbGORM struct {
	db *gorm.DB
}

func NewDBGORM(db *gorm.DB) *dbGORM {
	return &dbGORM{
		db: db,
	}
}

func (orm *dbGORM) FirstByPK(pk int, result any) error {
	return orm.db.First(
		result,
		pk,
	).
		Error
}

func (orm *dbGORM) GetProductByPK(pk int, result any) error {
	return orm.db.Raw(
		"select id, created_at, updated_at, deleted_at, code, price from products where id=$1",
		pk,
	).
		Scan(result).
		Error
}
