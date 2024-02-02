package ddltable

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Person struct {
	Persons struct{} `hera:"tablename"`

	ID              uint `hera:"pk"`
	Name            string
	Age             int16
	AllowedToDrive  bool `hera:"default:false, columnname:driving,"`
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

	fmt.Println(table.Columns)

	fmt.Println(table.AsDDLPostgres())
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
