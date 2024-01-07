package main

import "strings"

type state struct {
	rodFrom rod
	rodTo   rod
	rodAux  rod

	numDisks int
}

func newHanoi() *state {
	return &state{
		rodFrom:  rod{4, 3, 2, 1},
		numDisks: 4,
	}
}

func (s *state) isValid() bool {
	return s.rodFrom.isValid() && s.rodTo.isValid() && s.rodAux.isValid()
}

func (s state) String() string {
	var res []string

	align := func(content string) string {
		if len(content) == s.numDisks {
			return content
		}

		result := content

		for i := 0; i < s.numDisks-1; i++ {
			result = " " + result

			if len(result) == s.numDisks {
				return result
			}
		}

		return result
	}

	for i := s.numDisks - 1; i >= 0; i-- {
		res = append(res,
			align(s.rodFrom.getGraphRepresent(i))+
				"|"+
				s.rodTo.getGraphRepresent(i)+
				"|"+
				s.rodAux.getGraphRepresent(i),
		)
	}

	return strings.Join(res, "\n")
}
