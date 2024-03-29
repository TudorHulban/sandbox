package main

func twoSumOrdered(slice []int, target int) [2]int {
	i, j := 0, len(slice)-1

	for {
		if slice[i]+slice[j] == target {
			return [2]int{
				slice[i],
				slice[j],
			}
		}

		if slice[i]+slice[j] > target {
			j--

			continue
		}

		i++
	}
}
