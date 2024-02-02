package ddltable

func reflectToPG(reflectType string, isPK bool) (string, bool) {
	if isPK {
		switch reflectType {
		case "int", "int64", "uint":
			return "bigserial", false

		default:
			return "serial", false
		}
	}

	switch reflectType {
	case "string", "*string":
		return "text", false

	case "sql.NullString":
		return "text", true

	case "int8", "int16":
		return "smallint", false

	case "int32":
		return "integer", false

	case "int", "int64", "uint":
		return "bigint", false

	case "float64", "*float64":
		return "numeric", false

	case "bool", "*bool":
		return "boolean", false
	}

	return "", false
}
