package main

type structWOPointers struct {
	state string

	f1 string //nolint:unused
	f2 string //nolint:unused
	f3 string //nolint:unused
	f4 string //nolint:unused
}

type structWPointers struct {
	state *string

	f1 *string //nolint:unused
	f2 *string //nolint:unused
	f3 *string //nolint:unused
	f4 *string //nolint:unused
}

func getState1(p structWOPointers) string {
	return p.state
}

// worst performance
func getState2(p structWPointers) string {
	return *p.state
}

func getState3(p *structWOPointers) string {
	return p.state
}

func getState4(p *structWOPointers) string {
	return (*p).state
}

// performance is worst than passing by value
func getState5(p *structWPointers) string {
	return *p.state
}

// same performance as passing the struct by value
func getString(p string) string {
	return p
}

func main() {}
