package main

import "fmt"

func main() {
	firstName := "Kush"
	lastName := "Srigiri"
	dayOfBirth := 29
	monthOfBirth := 06
	yearOfBirth := 2007
	numberOfSiblings := 1
	heightInMeters := 1.70
	zodiacSign := '\u264b' // TODO: calculate using the provided formula!
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
