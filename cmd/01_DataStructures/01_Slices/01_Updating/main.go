package main

type person struct {
	name    string
	numbers []int
}

type persons []*person

func newPerson(name string, numbers ...int) *person {
	return &person{
		name:    name,
		numbers: numbers,
	}
}

func main() {}

func updateByValue(p []person, number int) {
	for _, person := range p {
		person.numbers = append(person.numbers, number)
	}
}

// updateByReference has better performance
func updateByReference(p *[]person, number int) {
	for i := range *p {
		(*p)[i].numbers = append((*p)[i].numbers, number)
	}
}

func updateByPointers(p []*person, number int) {
	for _, person := range p {
		person.numbers = append(person.numbers, number)
	}
}

// updatePreferred has better readability
func updatePreferred(p persons, number int) {
	for _, person := range p {
		person.numbers = append(person.numbers, number)
	}
}
