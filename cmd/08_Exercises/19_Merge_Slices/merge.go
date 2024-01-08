package main

func merge(ordered1, ordered2 []int) []int {
	var res []int

	appnd := func(s []int, pos int) {
		res = append(res, s[pos:]...)
	}

	var i, j int

	for {
		if i >= len(ordered1) {
			appnd(ordered2, j)

			break
		}

		if j >= len(ordered2) {
			appnd(ordered1, i)

			break
		}

		if ordered1[i] < ordered2[j] {
			res = append(res, ordered1[i])
			i++

			continue
		}

		res = append(res, ordered2[j])

		j++
	}

	return res
}
