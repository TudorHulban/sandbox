package main

func main() {}

func byLength(slice []string) bool {
	return len(slice) == 0
}

func byNil(slice []string) bool {
	return slice == nil
}
