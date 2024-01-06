package main

type Signal struct {
	Value int `json:"value"`
}

type Signals []Signal

var state = Signals{
	Signal{
		Value: 101,
	},

	Signal{
		Value: 78,
	},
}
