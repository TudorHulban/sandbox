package main

import (
	"errors"
	"fmt"
)

func main() {
	s := "abcdadba"

	fmt.Println(
		wins(extractOccurences(identifyLetters(s))),
	)
}

type winner string
type loser string

func wins(m []MapString) (winner, loser, error) {
	if len(m) == 0 {
		return "", "",
			errors.New("no votes")
	}

	if len(m) == 1 {
		return "", "",
			errors.New("no winner or loser")
	}

	if m[0].Value == m[len(m)-1].Value {
		return "", "",
			errors.New("no winner or loser")
	}

	if m[len(m)-1] == m[len(m)-2] {
		return "", loser(m[0].Key),
			errors.New("with looser but no winner")
	}

	if m[0].Value == m[1].Value {
		return winner(m[len(m)-1].Key), "",
			errors.New("with winner but no loser")
	}

	return winner(m[len(m)-1].Key), loser(m[0].Key),
		nil
}
