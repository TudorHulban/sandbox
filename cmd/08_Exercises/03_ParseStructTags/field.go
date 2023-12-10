package main

import (
	"fmt"
	"strings"
)

type Field struct {
	FieldName string
	Default   string // default value for field
	DataType  string
	FieldTag  string

	IsNullable bool
	Unique     bool
	Index      bool
}

type Fields []*Field

func (f Fields) String() string {
	var res []string

	for _, field := range f {
		res = append(res, fmt.Sprintf("%#v", *field))
	}

	return strings.Join(res, "\n")
}
