package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Code  string
	Price uint
}

func main() {
	db, errOpen := gorm.Open(
		sqlite.Open("test.dbf"),
		&gorm.Config{},
	)
	if errOpen != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&Product{},
	)

	db.Create(
		&Product{
			Code:  "D42",
			Price: 100,
		},
	)

	var product Product

	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")

	db.Model(&product).Update("Price", 200)
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	db.Delete(&product, 1)
}
