package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var paramsPG = ConfigPostgres{
	DBName:   "sandbox",
	User:     "postgres",
	Password: "password",
	Host:     "localhost",
	Port:     5432,
}

func main() {
	connGORM, errConnGORM := gorm.Open(
		postgres.Open(
			paramsPG.AsDSNGORM(),
		),
		&gorm.Config{
			DisableAutomaticPing: true,
		},
	)
	if errConnGORM != nil {
		fmt.Println(errConnGORM)

		os.Exit(10)
	}

	errMigration := connGORM.AutoMigrate(
		&Product{},
	)
	if errMigration != nil {
		fmt.Println(errMigration)

		os.Exit(11)
	}

	testItem := Product{
		Code:  "D43",
		Price: 100,
	}

	connGORM.Clauses(
		clause.OnConflict{DoNothing: true},
	).
		Create(&testItem)

	ctx := context.Background()

	connPGX, errConnPGX := pgx.Connect(
		ctx,
		paramsPG.AsDSNPGX(),
	)
	if errConnPGX != nil {
		fmt.Println(errConnPGX)

		os.Exit(12)
	}
	defer connPGX.Close(
		context.Background(),
	)

	var reconstructedProduct Product

	errGetReconstructed := connPGX.QueryRow(
		ctx,
		"select id, created_at, updated_at, deleted_at, code, price from products where id=$1",
		1,
	).
		Scan(
			&reconstructedProduct.ID,
			&reconstructedProduct.CreatedAt,
			&reconstructedProduct.UpdatedAt,
			&reconstructedProduct.DeletedAt,
			&reconstructedProduct.Code,
			&reconstructedProduct.Price,
		)
	if errGetReconstructed != nil {
		fmt.Println(errGetReconstructed)

		os.Exit(13)
	}

	fmt.Println(
		reconstructedProduct,
	)
}
