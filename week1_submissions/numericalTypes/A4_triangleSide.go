package main

import (
	"fmt"
	"math"
)

func main() {
	var side1, side2, angle float64

	fmt.Println("Enter the length of side 1: ")
	fmt.Scanln(&side1)

	fmt.Println("Enter the length of side 2: ")
	fmt.Scanln(&side2)

	fmt.Println("Enter the angle between the two sides: ")
	fmt.Scanln(&angle)

	side3 := math.Sqrt(math.Pow(side1, 2) + math.Pow(side2, 2) - 2 * side1 * side2 * math.Cos(angle * math.Pi / 180))

	fmt.Printf("The length of side 3 is: %.2f", side3)
}