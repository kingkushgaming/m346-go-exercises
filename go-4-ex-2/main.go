package main

import "math"

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	a := 3.0
	b := 4.0
	hypotenuse := computeHypotenuse(a, b)
	println("The length of the hypotenuse is", hypotenuse)
}
