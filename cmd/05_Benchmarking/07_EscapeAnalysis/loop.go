package escapeanalysis

// memory alignment reversed to make big memory consumtion
type car struct {
	isAWD  bool
	engine int

	price float32

	brand string
}

// memory alignment reversed to make big memory consumtion
type person struct {
	isMarried bool
	age       uint
	children  uint

	hasDog bool

	salary float32

	name    string
	surname string

	*car
}

func newCar() *car {
	return &car{
		isAWD:  true,
		engine: 2000,
		price:  1000,
		brand:  "abcd",
	}
}

func newPerson(ix int) *person {
	return &person{
		age:       uint(ix + 1),
		salary:    float32(ix),
		hasDog:    ix == 1,
		isMarried: ix == 1,
		children:  uint(ix),

		name:    "xyz",
		surname: "abc",

		car: newCar(),
	}
}

func ages(howMany uint) []person {
	var res []person

	for i := 0; i < int(howMany); i++ {
		res = append(res, *newPerson(i))
	}

	return res
}

func agesPtr(howMany uint) []*person {
	var res []*person

	for i := 0; i < int(howMany); i++ {
		res = append(res, &person{
			age: uint(i + 1),

			car: newCar(),
		})
	}

	return res
}

func agesPtrOptim(howMany uint) []*person {
	res := make([]*person, howMany)

	for i := 0; i < int(howMany); i++ {
		res[i] = &person{
			age: uint(i + 1),

			car: newCar(),
		}
	}

	return res
}

func agesResults(howMany uint, results *[]*person) {
	for i := 0; i < int(howMany); i++ {
		*results = append(*results, &person{
			age: uint(i + 1),

			car: newCar(),
		})
	}
}

func agesResultsOptim(howMany uint, results *[]*person) {
	for i := 0; i < int(howMany); i++ {
		(*results)[i] = &person{
			age: uint(i + 1),

			car: newCar(),
		}
	}
}
