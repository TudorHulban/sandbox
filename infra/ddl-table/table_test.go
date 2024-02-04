package ddltable

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Person struct {
	Persons struct{} `hera:"tablename"`

	Name string

	ID uint `hera:"pk, order:0"`

	Age             int16
	AllowedToDrive  bool `hera:"default:false, columnname:driving, order:2"`
	skipNotExported bool
	SkipExported    bool `hera:"-"`
	Birthdate       sql.NullString
}

type PersonsInGroups struct {
	IDPersons uint `hera:"index:ix_personsingroups"`
	IDGroups  uint `hera:"index:ix_personsingroups"`
}

func TestPersonsTable(t *testing.T) {
	table, errNew := NewTable(
		&Person{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	// fmt.Println(table.Columns)

	fmt.Println(table.MigrationTable)
	fmt.Println(table.MigrationIndexes)
}

func TestTablePersonsInGroups(t *testing.T) {
	table, errNew := NewTable(
		&PersonsInGroups{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	// fmt.Println(table.Columns)

	fmt.Println(table.MigrationTable)
	fmt.Println(table.MigrationIndexes)
}
