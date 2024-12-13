package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}
type BirthDate struct {
	Day   byte
	Month string
	Year  uint16
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "John",
			LastName:  "Doe",
		},
		BirthDate: BirthDate{
			Day:   12,
			Month: "December",
			Year:  1990,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       'S',
	}
	fmt.Println(me)

	me.NumberOfSiblings++
	fmt.Println("Siblings Before:", me.NumberOfSiblings-1)
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
