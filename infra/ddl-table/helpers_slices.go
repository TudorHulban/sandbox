package ddltable

func High[T comparable](slice []T) T {
	return slice[len(slice)-1]
}

func Pop[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}

	return append(slice[:0], slice[:len(slice)-1]...)
}
