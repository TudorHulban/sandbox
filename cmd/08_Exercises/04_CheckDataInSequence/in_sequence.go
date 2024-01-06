package main

import "errors"

func inSequence(data, sequence []int) error {
	if len(sequence) == 0 || len(data) < len(sequence) {
		return errInputValidation
	}

	sequenceLength := len(sequence)

	for i := range data {
		if i%sequenceLength == 0 {
			for ix, pos := range sequence {
				if i+ix > len(data)-1 {
					return errors.New("not enough values for sequence")
				}

				if data[i+ix] == pos {
					continue
				}

				return errors.New("not in sequence")
			}
		}
	}

	return nil
}
