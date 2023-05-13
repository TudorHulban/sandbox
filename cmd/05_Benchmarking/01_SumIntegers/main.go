package main

func main() {}

func sumLoop(ints []int) int {
	if len(ints) == 0 {
		return 0
	}

	var res int
	for _, v := range ints {
		res = res + v
	}

	return res
}

func sumRecurs(ints []int) int {
	if len(ints) == 0 {
		return 0
	}

	return ints[0] + sumRecurs(ints[1:])
}
