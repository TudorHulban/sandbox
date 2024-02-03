package ddltable

import (
	"fmt"
	"strings"
)

type Table struct {
	Name    string
	Columns columns
	Indexes map[string]*index
}

func NewTable(object any) (*Table, error) {
	columns, overrideTableName, errGetColumns := newColumns(object)
	if errGetColumns != nil {
		return nil,
			errGetColumns
	}

	columns.sortForColumnOrder()

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
			Indexes: make(map[string]*index),
		},
		nil
}

// AsDDLPostgres returns DDL Table, DDL Table index(es).
func (t *Table) AsDDLPostgres() (string, string, error) {
	ddlTable := t.ddlTable()
	ddlIndex, errDDLIndex := t.ddlIndex()
	if errDDLIndex != nil {
		return "", "",
			errDDLIndex
	}

	if len(ddlIndex) == 0 {
		return ddlTable,
			"", nil
	}

	return ddlTable, ddlIndex,
		nil
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

func (t *Table) ddlIndex() (string, error) {
	var indexName string

	for _, column := range t.Columns {
		if column.IsIndexed {
			if len(column.IndexName) == 0 {
				indexName = t.Name + "_idx"
			} else {
				indexName = column.IndexName
			}

			if _, exists := t.Indexes[column.IndexName]; exists {
				if column.IndexType != t.Indexes[column.IndexName].Type {
					return "",
						fmt.Errorf(
							"for column %s, index type (%s) is different than previous index type for index name %s",
							column.Name,
							column.IndexType,
							indexName,
						)
				}

				t.Indexes[indexName].ColumnNames = append(t.Indexes[indexName].ColumnNames, column.Name)

				continue
			}

			t.Indexes[indexName] = &index{
				Type: column.IndexType,
				ColumnNames: []string{
					column.Name,
				},
			}
		}
	}

	if len(t.Indexes) == 0 {
		return "", nil
	}

	return t.renderIndexes(),
		nil
}

func (t *Table) renderIndexes() string {
	result := make([]string, 0)

	for indexName, indexInfo := range t.Indexes {
		result = append(result,
			indexInfo.render()(t.Name, indexName),
		)
	}

	return strings.Join(result, "\n")
}
