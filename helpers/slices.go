package helpers

func High[T comparable](slice []T) T {
	return slice[len(slice)-1]
}

func Low[T comparable](slice []T) T {
	return slice[0]
}
