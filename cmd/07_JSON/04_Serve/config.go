package main

type FieldConfig struct {
	Name   string `json:"name"`
	Type   int    `json:"ftype"`
	Length int    `json:"flength"`

	MinValue int `json:"min"`
	MaxValue int `json:"max"`

	PositiveOnly bool `json:"fpositive"`
}

type GenConfig struct {
	Configuration []FieldConfig `json:"config"`
	NoRows        int64         `json:"norows"`
}

type GenData struct {
	ColumnNames []string        `json:"columns"`
	Rows        [][]interface{} `json:"rowsdata"`
}

func NewConfig(noRows int64) *GenConfig {
	instance := GenConfig{
		NoRows:        noRows,
		Configuration: make([]FieldConfig, 2),
	}

	instance.Configuration[0] = FieldConfig{
		Name:         "a",
		Type:         2,
		Length:       5,
		MinValue:     0,
		MaxValue:     100,
		PositiveOnly: true,
	}

	instance.Configuration[0] = FieldConfig{
		Name:         "b",
		Type:         2,
		Length:       5,
		MinValue:     0,
		MaxValue:     100,
		PositiveOnly: true,
	}

	return &instance
}
