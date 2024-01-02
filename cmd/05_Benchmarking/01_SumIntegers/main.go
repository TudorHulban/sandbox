package main

func main() {}

func sumLoop(values []int) int {
	if len(values) == 0 {
		return 0
	}

	var result int

	for _, value := range values {
		result = result + value
	}

	return result
}

func sumRecurs(values []int) int {
	if len(values) == 0 {
		return 0
	}

	return values[0] + sumRecurs(values[1:])
}
