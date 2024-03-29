package main

import (
	"errors"
	"log"
	"strconv"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) > 0 {
		result := (*s)[len(*s)-1]

		*s = (*s)[0 : len(*s)-1]

		return result
	}

	return 0
}

func getSqDist(p1, p2 []float64) (float64, error) {
	if (len(p1) != 2) && (len(p2) != 2) {
		return 0, errors.New("Problems with points: " + "p1: " + strconv.FormatInt(int64(len(p1)), 10) + "p2: " + strconv.FormatInt(int64(len(p2)), 10))
	}

	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]

	return dx*dx + dy*dy, nil
}

func getSqSegDist(p, p1, p2 []float64) float64 {
	x := p1[0]
	y := p1[1]

	dx := p2[0] - x
	dy := p2[1] - y

	if dx != 0 || dy != 0 {
		t := ((p[0]-x)*dx + (p[1]-y)*dy) / (dx*dx + dy*dy)
		if t > 1 {
			x = p2[0]
			y = p2[1]
		} else if t > 0 {
			x += dx * t
			y += dy * t
		}
	}

	dx = p[0] - x
	dy = p[1] - y

	return dx*dx + dy*dy
}

func simplifyRadialDist(points [][]float64, sqTolerance float64) [][]float64 {
	prevPoint := points[0]
	newPoints := [][]float64{prevPoint}

	var point []float64

	for i := 1; i < len(points); i++ {
		point = points[i]

		delta, err := getSqDist(point, prevPoint)
		if err != nil {
			log.Panicln("getSqDist")
		}

		if delta > sqTolerance {
			newPoints = append(newPoints, point)

			prevPoint = point
		}
	}

	if !ComparePoints(prevPoint, point) {
		newPoints = append(newPoints, point)
	}

	return newPoints
}

func simplifyDouglasPeucker(points [][]float64, sqTolerance float64) [][]float64 {
	var l = len(points)
	markers := make([]int, l)

	first := 0
	last := l - 1

	var stack Stack
	var newPoints [][]float64

	i, index := 0, 0
	maxSqDist, sqDist := float64(0), float64(0)
	markers[first], markers[last] = 1, 1

	for last > 0 {
		maxSqDist = 0

		for i = first + 1; i < last; i++ {
			sqDist = getSqSegDist(points[i], points[first], points[last])
			if sqDist > maxSqDist {
				index = i
				maxSqDist = sqDist
			}
		}

		if maxSqDist > sqTolerance {
			markers[index] = 1

			stack.Push(first)
			stack.Push(index)
			stack.Push(index)
			stack.Push(last)
		}

		last = stack.Pop()
		first = stack.Pop()
	}

	for i = 0; i < l; i++ {
		if checkArrIndex(markers, i) {
			newPoints = append(newPoints, points[i])
		}
	}

	return newPoints
}

func checkArrIndex(arr []int, index int) bool {
	if index < len(arr) && index >= 0 {
		if arr[index] > 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Simplify(points [][]float64, tolerance float64, highestQuality bool) [][]float64 {
	if len(points) == 0 {
		return nil
	}

	sqTolerance := tolerance * tolerance

	var result [][]float64

	if highestQuality {
		result = points
	} else {
		result = simplifyRadialDist(points, sqTolerance)
	}

	result = simplifyDouglasPeucker(result, sqTolerance)

	return result
}

func CompareSlices(p1, p2 [][]float64) bool {
	if len(p1) == len(p2) {
		for i := range p1 {
			if !ComparePoints(p1[i], p2[i]) {
				return false
			}
		}

		return true
	}

	return false
}

func ComparePoints(p1, p2 []float64) bool {
	if p1[0] == p2[0] && p1[1] == p2[1] {
		return true
	}

	return false
}
