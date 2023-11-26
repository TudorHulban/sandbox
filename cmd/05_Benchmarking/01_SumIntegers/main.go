package main

func main() {}

func sumLoop(values []int) int {
	if len(values) == 0 {
		return 0
	}

	var res int
	for _, value := range values {
		res = res + value
	}

	return res
}

func sumRecurs(values []int) int {
	if len(values) == 0 {
		return 0
	}

	return values[0] + sumRecurs(values[1:])
}
