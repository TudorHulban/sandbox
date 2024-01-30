package ddltable

import (
	"fmt"
	"strings"
)

type columns []*column

func (cols columns) String() string {
	result := []string{
		fmt.Sprintf(
			"Columns: %d",
			len(cols),
		),
	}

	for ix, column := range cols {
		result = append(result,
			fmt.Sprintf(
				"%d. Name: %s, Type: %s, IsPK: %t, IsNullable: %t, IsIndexed: %t, IsToBeSkipped: %t, IndexName: %s",
				ix+1,
				column.Name,
				column.PGType,
				column.IsPK,
				column.IsNullable,
				column.IsIndexed,
				column.IsToBeSkipped,
				column.IndexName,
			),
		)
	}

	return strings.Join(result, "\n")
}
