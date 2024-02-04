package ddltable

import (
	"fmt"
	"testing"
)

func TestIndexRender(t *testing.T) {
	ix := index{
		ColumnNames: []string{"a", "b"},
	}

	fmt.Println(ix.migrationUp()("table-name", "index-name"))
}
