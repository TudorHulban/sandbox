package main

type rod []int

func (r rod) getLargestDisk() int {
	if len(r) == 0 {
		return 0
	}

	return r[0]
}

func (r rod) getSmallestDisk() int {
	if len(r) == 0 {
		return 0
	}

	return r[len(r)-1]
}

func (r rod) isValid() bool {
	if len(r) <= 1 {
		return true
	}

	for i := 0; i < len(r)-1; i++ {
		if r[i] < r[i+1] {
			return false
		}
	}

	return true
}

func (r rod) getGraphRepresent(position int) string {
	if len(r) == 0 {
		return ""
	}

	getX := func(howMany int) string {
		var result string

		for i := 0; i < howMany; i++ {
			result = result + "x"
		}

		return result
	}

	return getX(r[position])
}
