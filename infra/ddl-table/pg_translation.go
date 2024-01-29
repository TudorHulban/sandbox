package ddltable

func reflectToPG(reflectType string, isPK bool) string {
	if isPK {
		switch reflectType {
		case "int", "int64", "uint":
			return "bigserial"

		default:
			return "serial"
		}
	}

	switch reflectType {
	case "string", "*string":
		return "text"

	case "int8", "int16":
		return "smallint"

	case "int32":
		return "integer"

	case "int", "int64", "uint":
		return "bigint"

	case "float64", "*float64":
		return "numeric"

	case "bool", "*bool":
		return "boolean"
	}

	return ""
}
