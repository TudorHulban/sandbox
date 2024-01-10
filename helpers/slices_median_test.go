package helpers

import "testing"

func TestMedian(t *testing.T) {
	cases := []struct {
		description string
		input       []int
		expected    float64
	}{
		{
			"1. zero",
			[]int{0},
			0.,
		},
		{
			"2. one",
			[]int{1},
			1.,
		},
		{
			"3. two",
			[]int{2},
			2.,
		},
		{
			"4. odd length",
			[]int{1, 2, 3},
			2.,
		},
		{
			"5. even length",
			[]int{1, 2, 3, 4},
			2.5,
		},
	}

	for _, tc := range cases {
		m := Median(tc.input)

		if m != tc.expected {
			t.Log(tc.description, m)

			t.FailNow()
		}
	}
}
