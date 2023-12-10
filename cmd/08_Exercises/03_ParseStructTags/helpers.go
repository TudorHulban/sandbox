package main

import (
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

func getTableName(model any) string {
	if typeOfModel := reflect.TypeOf(model); typeOfModel.Kind() == reflect.Ptr {
		return inflection.Plural(
			strcase.ToSnake(typeOfModel.Elem().Name()),
		)
	}

	return inflection.Plural(
		strcase.ToSnake(reflect.TypeOf(model).Name()),
	)
}

func getAllowedTypes() []string {
	result := []string{"string", "int", "int64", "time.Time", "float64", "bool"}

	for _, allowedType := range result {
		result = append(result, "*"+allowedType)
		result = append(result, "[]"+allowedType)
		result = append(result, "[]*"+allowedType)
	}

	result = append(result, "map[string]interface")
	result = append(result, "interface")

	return result
}

func getFields(model any) Fields {
	allowedTypes := getAllowedTypes()

	isAllowedType := func(fieldType string) bool {
		for _, at := range allowedTypes {
			if fieldType == at {
				return true
			}
		}

		return false
	}

	val := reflect.ValueOf(model).Elem()

	var setFieldType string // takes field type from tag instead

	var result []*Field

	for i := 0; i < val.NumField(); i++ {
		field := Field{
			FieldName: strcase.ToSnake(val.Type().Field(i).Name),
		}

		fieldType := val.Type().Field(i).Type.String()
		tag := val.Type().Field(i).Tag.Get(_tagName)

		if len(tag) > 0 {
			tags := strings.Split(tag, ",")

			for _, tagS := range tags {
				s := strings.ToLower(strings.TrimSpace(tagS))

				if s == "-" {
					continue
				}

				if s == "unique" {
					field.Unique = true
				}

				if s == "index" {
					field.Index = true
				}

				ss := strings.Split(s, "type:")
				if len(ss) > 1 {
					field.DataType = ss[1]
					setFieldType = ss[1]
				}
			}
		}

		field.FieldTag = tag

		if !isAllowedType(fieldType) {
			continue
		}

		if fieldType == "string" {
			field.DataType = "text"
			field.IsNullable = false
			field.Default = "''"
		}

		if fieldType == "*string" {
			field.DataType = "text"
			field.IsNullable = true
		}

		if fieldType == "[]string" || fieldType == "[]*string" {
			field.DataType = "text[]"
		}

		if fieldType == "int64" || fieldType == "int" {
			field.DataType = "bigint"
			field.IsNullable = false
		}

		if fieldType == "*int64" {
			field.DataType = "bigint"
			field.IsNullable = true
		}

		if fieldType == "[]int64" || fieldType == "[]*int64" {
			field.DataType = "integer[]"
			field.IsNullable = true
		}

		if fieldType == "*time.Time" {
			field.DataType = "timestamptz"
			field.IsNullable = true
		}

		if fieldType == "time.Time" {
			field.IsNullable = false
			field.DataType = "timestamptz"
			field.Default = "NOW()"
		}

		if fieldType == "float64" {
			field.DataType = "numeric"
			field.IsNullable = false
			field.Default = "0.00"
		}

		if fieldType == "*float64" {
			field.DataType = "numeric"
			field.IsNullable = true
		}

		if fieldType == "[]float64" || fieldType == "[]*float64" {
			field.DataType = "numeric[]"
			field.IsNullable = true
		}

		if fieldType == "bool" {
			field.DataType = "boolean"
			field.IsNullable = false
			field.Default = "false"
		}

		if fieldType == "*bool" {
			field.DataType = "boolean"
			field.IsNullable = true
		}

		if fieldType == "[]bool" || fieldType == "[]*bool" {
			field.DataType = "boolean[]"
			field.IsNullable = true
		}

		if fieldType == "map[string]interface" || fieldType == "interface" {
			field.DataType = "jsonb"
			field.IsNullable = true
		}

		if len(setFieldType) > 0 {
			field.DataType = setFieldType
		}

		result = append(result, &field)
	}

	return result
}
