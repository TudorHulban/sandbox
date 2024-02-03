package ddltable

import (
	"fmt"
	"strings"
)

type index struct {
	Type        string
	ColumnNames []string
}

func (ix index) render() func(nameTable, nameIndex string) string {
	return func(nameTable, nameIndex string) string {
		return fmt.Sprintf(
			"create index %s if not exists %s on %s ` ("+strings.Join(ix.ColumnNames, ",")+");`",
			ix.Type,
			nameIndex,
			nameTable,
		)
	}
}
