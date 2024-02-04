package ddltable

import (
	"fmt"
	"strings"
)

type index struct {
	Type        string
	ColumnNames []string
}

func (ix index) migrationUp() func(nameTable, nameIndex string) string {
	return func(nameTable, nameIndex string) string {
		return fmt.Sprintf(
			"create index %s if not exists %s on %s("+strings.Join(ix.ColumnNames, ",")+");",
			ix.Type,
			nameIndex,
			nameTable,
		)
	}
}

func (ix index) migrationDown() func(nameIndex string) string {
	return func(nameIndex string) string {
		return fmt.Sprintf(
			"drop index if exists %s;",
			nameIndex,
		)
	}
}
