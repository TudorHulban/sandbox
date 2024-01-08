package main

import (
	"fmt"
	"sort"
)

func mergeBrute(a ...[]int) []int {
	if len(a) == 0 {
		return []int{}
	}

	buf := []elem{}

	for i, v := range a {
		buf = append(buf, elem{
			val: v[0],
			pos: i,
		})
	}

	sort.Slice(buf, func(i, j int) bool {
		return buf[i].val < buf[j].val
	})

	var res []int

	for _, v := range buf {
		res = append(res, a[v.pos]...)
	}

	fmt.Println(res)

	sort.Ints(res)

	return res
}
