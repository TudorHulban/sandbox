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

	fmt.Println(table.Columns)

	fmt.Println(table.AsDDLPostgres())
}

func TestTablePersonsInGroups(t *testing.T) {
	table, errNew := NewTable(
		&PersonsInGroups{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	// fmt.Println(table.Columns)

	ddlTable, ddlIndexes, errDDL := table.AsDDLPostgres()
	require.NoError(t, errDDL)

	fmt.Println(ddlTable)
	fmt.Println(ddlIndexes)
}

func BenchmarkTable(b *testing.B) {
	table, errNew := NewTable(
		&Person{},
	)
	require.NoError(b, errNew)
	require.NotZero(b, table)

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		table.AsDDLPostgres()
	}
}
