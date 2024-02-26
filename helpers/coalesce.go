package helpers

func Coalesce[T comparable](item1, item2 T) T {
	var zero T

	if item1 == zero {
		return item2
	}

	return item1
}
