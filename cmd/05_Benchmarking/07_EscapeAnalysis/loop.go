package escapeanalysis

type carReversedAlignment struct { //nolint:govet
	isAWD  bool
	engine int
	price  float32
	brand  string
}

type personReversedAlignment struct { //nolint:govet
	isMarried bool
	hasDog    bool

	age      uint
	children uint

	*carReversedAlignment

	salary float32

	name    string
	surname string
}

func newCar() *carReversedAlignment {
	return &carReversedAlignment{
		isAWD:  true,
		engine: 2000,
		price:  1000,
		brand:  "abcd",
	}
}

func newPerson(ix int) *personReversedAlignment {
	return &personReversedAlignment{
		age:       uint(ix + 1),
		salary:    float32(ix),
		hasDog:    ix == 1,
		isMarried: ix == 1,
		children:  uint(ix),

		name:    "xyz",
		surname: "abc",

		carReversedAlignment: newCar(),
	}
}

func ages(howMany uint) []personReversedAlignment {
	var res []personReversedAlignment

	for i := 0; i < int(howMany); i++ {
		res = append(res, *newPerson(i))
	}

	return res
}

func agesPtr(howMany uint) []*personReversedAlignment {
	var res []*personReversedAlignment

	for i := 0; i < int(howMany); i++ {
		res = append(res, &personReversedAlignment{
			age: uint(i + 1),

			carReversedAlignment: newCar(),
		})
	}

	return res
}

func agesPtrOptim(howMany uint) []*personReversedAlignment {
	res := make([]*personReversedAlignment, howMany)

	for i := 0; i < int(howMany); i++ {
		res[i] = &personReversedAlignment{
			age: uint(i + 1),

			carReversedAlignment: newCar(),
		}
	}

	return res
}

func agesResults(howMany uint, results *[]*personReversedAlignment) {
	for i := 0; i < int(howMany); i++ {
		*results = append(*results, &personReversedAlignment{
			age: uint(i + 1),

			carReversedAlignment: newCar(),
		})
	}
}

func agesResultsOptim(howMany uint, results *[]*personReversedAlignment) {
	for i := 0; i < int(howMany); i++ {
		(*results)[i] = &personReversedAlignment{
			age: uint(i + 1),

			carReversedAlignment: newCar(),
		}
	}
}
