package ddltable

import (
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
}

func TestTable(t *testing.T) {
	table, errNew := NewTable(
		&Person{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	// fmt.Println(table.Columns)

	fmt.Println(table.AsDDLPostgres())
}
