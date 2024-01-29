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
				"%d. Name: %s, Type: %s, IsPK: %t, IsNullable: %t, IsIndexed: %t, IsToBeSkipped: %t",
				ix+1,
				column.Name,
				column.PGType,
				column.IsPK,
				column.IsNullable,
				column.IsIndexed,
				column.IsToBeSkipped,
			),
		)
	}

	return strings.Join(result, "\n")
}
