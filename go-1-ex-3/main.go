package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// Open files for writing
	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	diceLogFile, err := os.OpenFile("dice.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening dice.log:", err)
		return
	}
	defer diceLogFile.Close()

	// Use fmt.Fprintln to print and write to files
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")
	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")

	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)
	fmt.Fprintln(diceLogFile, "the dice was rolled at", when)
}
