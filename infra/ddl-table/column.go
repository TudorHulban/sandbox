package ddltable

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type column struct {
	Name         string
	PGType       string
	DefaultValue string
	IndexName    string

	IsPK          bool
	IsNullable    bool
	IsUnique      bool
	IsIndexed     bool
	IsToBeSkipped bool
}

func newColumns(object any) (columns, string, error) {
	result := make([]*column, 0)

	var alreadyHavePK bool

	var tableName string

	for i := 0; i < reflect.TypeOf(object).Elem().NumField(); i++ {
		fieldRoot := reflect.TypeOf(object).
			Elem().
			FieldByIndex([]int{i})

		column := column{
			Name: fieldRoot.Name,
		}

		if valueTag, hasTag := fieldRoot.Tag.Lookup(_TagName); hasTag {
			errUpdate := column.UpdateWith(valueTag, alreadyHavePK)
			if errUpdate != nil {
				if errUpdate.Error() == errIsOverrideTableName.Error() {
					tableName = strings.ToLower(fieldRoot.Name)

					continue
				}

				return nil, "",
					errUpdate
			}
		}

		if !fieldRoot.IsExported() {
			continue
		}

		column.PGType = reflectToPG(fieldRoot.Type.String(), column.IsPK)

		if column.IsPK {
			alreadyHavePK = true
		}

		result = append(result, &column)
	}

	return result,
		tableName,
		nil
}

func (col *column) UpdateWith(tagValues string, alreadyHavePK bool) error {
	for _, tagValue := range strings.Split(
		tagValues, ",",
	) {
		tagClean := strings.ToLower(
			strings.TrimSpace(tagValue),
		)

		var compoundTagValue string

		if strings.Contains(tagClean, _TagSeparator) {
			tagCompound := strings.Split(tagClean, _TagSeparator)

			if len(tagCompound) != 2 {
				return fmt.Errorf(
					"malformed tag value: %s",
					tagClean,
				)
			}

			tagClean = tagCompound[0]
			compoundTagValue = tagCompound[1]
		}

		if len(tagClean) == 0 {
			return nil
		}

		if tagClean == "-" {
			col.IsToBeSkipped = true

			return nil
		}

		if tagClean == _TagOverrideTableName {
			return errIsOverrideTableName
		}

		if tagClean == _TagPK {
			if alreadyHavePK {
				return errors.New("more than one primary key field detected. max is 1")
			}

			col.IsPK = true
		}

		if tagClean == _TagUnique {
			col.IsUnique = true
		}

		if tagClean == _TagIndexed {
			col.IsIndexed = true
			col.IndexName = compoundTagValue
		}

		if tagClean == _TagRequired {
			col.IsNullable = false
		}

		if tagClean == _TagDefault {
			col.DefaultValue = compoundTagValue
		}

		if tagClean == _TagOverrideColumnName {
			col.Name = compoundTagValue
		}
	}

	return nil
}

func (col *column) AsDDLPostgres() string {
	if col.IsToBeSkipped {
		return ""
	}

	result := []string{
		strings.ToLower(col.Name),
	}

	result = append(result,
		col.PGType,
	)

	if col.IsPK {
		result = append(result,
			"PRIMARY KEY",
		)
	}

	if col.IsUnique {
		result = append(result, "UNIQUE")
	}

	if col.IsNullable {
		result = append(result, "NOT NULL")
	}

	if len(col.DefaultValue) > 0 {
		result = append(result, "DEFAULT")
		result = append(result, col.DefaultValue)
	}

	return strings.Join(result, " ")
}
