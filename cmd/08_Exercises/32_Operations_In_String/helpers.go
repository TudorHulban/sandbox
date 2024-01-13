package main

func splitASCIIString(expression string, splitters []string) []string {
	var buf string

	var result []string

walk:
	for i := 0; i < len(expression); i++ {
		currentChar := expression[i : i+1]

		for _, splitter := range splitters {
			if currentChar == splitter {
				result = append(result, buf)
				buf = ""

				result = append(result, splitter)

				continue walk
			}
		}

		buf = buf + currentChar

		if i == len(expression)-1 {
			result = append(result, buf)
		}
	}

	return result
}
