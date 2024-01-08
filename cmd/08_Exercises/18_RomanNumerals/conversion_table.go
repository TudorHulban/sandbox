package main

var conversionMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "LC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var conversionValues = []uint{
	1,
	4,
	5,
	9,
	10,
	40,
	50,
	90,
	100,
	400,
	500,
	900,
	1000,
}
