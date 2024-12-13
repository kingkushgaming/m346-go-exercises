package main

import (
	"fmt"
	"os"
)

// TODO: implement the function computeGrade
func computeGrade(points float32, maxpoints float32) {
	if points < 0 || maxpoints < 0 {
		fmt.Fprintln(os.Stderr, "Points have to be positive")
		return
	}
	if points > maxpoints {
		fmt.Fprintln(os.Stderr, "Points cannot be bigger than max points")
		return
	}

	note := (points/maxpoints)*5 + 1
	fmt.Println(note)
}

func main() {
	// TODO: call the function computeGrade
	computeGrade(13.5, 20.0) // 4.375
	computeGrade(20, 10)     // Points cannot be bigger than max points
	computeGrade(-3, 20)     // Points have to be positive
}
