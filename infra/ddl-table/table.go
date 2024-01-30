package ddltable

import (
	"strings"
)

type Table struct {
	Name    string
	Columns columns
}

func NewTable(object any) (*Table, error) {
	columns, overrideTableName, errGetColumns := newColumns(object)
	if errGetColumns != nil {
		return nil,
			errGetColumns
	}

	var tableName string

	if len(overrideTableName) > 0 {
		tableName = overrideTableName
	} else {
		tableName = strings.ToLower(getObjectName(object))
	}

	if strings.HasPrefix(tableName, _TagPointer) {
		tableName = tableName[1:]
	}

	return &Table{
			Name:    pluralize(tableName),
			Columns: columns,
		},
		nil
}

func (t *Table) AsDDLPostgres() string {
	ddlTable := t.ddlTable()
	ddlIndex := t.ddlIndex()

	if len(ddlIndex) == 0 {
		return ddlTable
	}

	return strings.Join(
		[]string{
			ddlTable,
			ddlIndex,
		},
		"\n",
	)
}

func (t *Table) ddlTable() string {
	result := []string{
		"create table if not exists ",
		t.Name,
		"(",
	}

	for ix, column := range t.Columns {
		ddl := column.AsDDLPostgres()

		if len(ddl) == 0 {
			if ix == len(t.Columns)-1 {
				if High[string](result) == _FieldSeparator {
					result = Pop[string](result)
				}
			}

			continue
		}

		result = append(result,
			ddl,
		)

		if ix < len(t.Columns)-1 {
			result = append(result,
				_FieldSeparator,
			)
		}
	}

	result = append(result, ");")

	return strings.Join(result, "")
}

func (t *Table) ddlIndex() string {
	var indexName string

	var indexedColumns []string

	for _, column := range t.Columns {
		if column.IsIndexed {
			indexedColumns = append(indexedColumns,
				column.Name,
			)

			indexName = column.IndexName
		}
	}

	if len(indexedColumns) == 0 {
		return ""
	}

	if len(indexName) == 0 {
		indexName = t.Name + "_idx"
	}

	return "create index " + indexName + " ON " + t.Name + " (" + strings.Join(indexedColumns, ",") + ");"
}
