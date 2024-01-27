package main

import (
	"github.com/golang-migrate/migrate/v4"
)

func runDBMigration(migrationURL string, dbSource string) error {
	migration, errNewMigration := migrate.New(migrationURL, dbSource)
	if errNewMigration != nil {
		return errNewMigration
	}

	if errMigrationUp := migration.Up(); errMigrationUp != nil && errMigrationUp != migrate.ErrNoChange {
		return errMigrationUp
	}

	return nil
}
