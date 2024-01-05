package main

import (
	"errors"
	"strconv"
)

// SplitPointsSlice - returns sliceLabels, sliceData
func SplitPointsSlice(slice [][]float64) ([]float64, []float64, error) {
	// assumption - input is slice of Points, aka 2 values
	sliceLabels := []float64{}
	sliceData := []float64{}

	for k, v := range slice {
		if len(v) != 2 {
			return nil, nil,
				errors.New("slice has issues at row: " + strconv.FormatInt(int64(k), 10))
		}

		sliceLabels = append(sliceLabels, v[0])
		sliceData = append(sliceData, v[1])
	}

	return sliceLabels, sliceData, nil
}

func SliceInterfToFloat(slice *[][]interface{}) (*[][]float64, error) {
	instance := make([][]float64, 0)

	for _, v := range *slice {
		row := []float64{}

		for _, vv := range v {
			switch vv.(type) {
			case float64:
				row = append(row, vv.(float64))
			case int:
				row = append(row, float64(vv.(int)))
			}
		}
		instance = append(instance, row)
	}
	return &instance, nil
}
