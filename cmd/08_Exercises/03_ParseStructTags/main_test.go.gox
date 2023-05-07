package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTableName(t *testing.T) {
	require.Equal(t, "users", getTableName(&User{}))
}

func TestGetFields(t *testing.T) {
	fields := getFields(&User{})

	require.Equal(t, "name", fields[0].FieldName)
	require.Equal(t, "age", fields[1].FieldName)

	require.Equal(t, "fieldName1", fields[0].FieldTag)
	require.Equal(t, "fieldName2", fields[1].FieldTag)
}
