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

	case "sql.NullInt16":
		return "smallint", true

	case "int32":
		return "integer", false

	case "sql.NullInt32":
		return "integer", true

	case "int", "int64", "uint":
		return "bigint", false

	case "sql.NullInt64":
		return "bigint", true

	case "float64", "*float64":
		return "DOUBLE PRECISION", false

	case "sql.NullFloat64":
		return "DOUBLE PRECISION", true

	case "bool", "*bool":
		return "boolean", false
	}

	return "", false
}
