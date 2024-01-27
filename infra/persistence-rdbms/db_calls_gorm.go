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

func (orm *dbGORM) FirstByPK(pk int, result any) {
	orm.db.First(
		result,
		pk,
	)
}
