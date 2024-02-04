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
func (t *Table) AsDDLPostgres() (*Migration, *Migration, error) {
	ddlIndexMigrationUp, ddlIndexMigrationDown, errDDLIndex := t.ddlIndexMigrationUp()
	if errDDLIndex != nil {
		return nil, nil,
			errDDLIndex
	}

	if len(ddlIndexMigrationUp) == 0 && len(ddlIndexMigrationDown) == 0 {
		return &Migration{
				Up:   t.ddlTableMigrationUp(),
				Down: t.ddlTableMigrationDown(),
			},
			nil, nil
	}

	return &Migration{
			Up:   t.ddlTableMigrationUp(),
			Down: t.ddlTableMigrationDown(),
		},
		&Migration{
			Up:   ddlIndexMigrationUp,
			Down: ddlIndexMigrationDown,
		},
		nil
}

func (t *Table) ddlTableMigrationUp() string {
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

func (t *Table) ddlTableMigrationDown() string {
	return fmt.Sprintf(
		"drop table if exists %s",
		t.Name,
	)
}

func (t *Table) ddlIndexMigrationUp() (string, string, error) {
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
					return "", "",
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
		return "", "", nil
	}

	indexesMigrationUp, indexesMigrationDown := t.renderIndexes()

	return indexesMigrationUp, indexesMigrationDown,
		nil
}

func (t *Table) renderIndexes() (string, string) {
	resultUp, resultDown := make([]string, 0), make([]string, 0)

	for indexName, indexInfo := range t.Indexes {
		resultUp = append(resultUp,
			indexInfo.migrationUp()(t.Name, indexName),
		)

		resultDown = append(resultDown,
			indexInfo.migrationDown()(t.Name),
		)
	}

	return strings.Join(resultUp, "\n"), strings.Join(resultDown, "\n")
}
