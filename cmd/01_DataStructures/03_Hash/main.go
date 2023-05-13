package main

import "fmt"

type state map[string]int

func newState() state {
	return make(map[string]int)
}

func (s state) addEntry(state string, value int) {
	if _, exists := s[state]; exists {
		return
	}

	s[state] = value
}

func main() {
	state := newState()
	state.addEntry("1", 1)

	fmt.Println(state) // map[1:1]
}
