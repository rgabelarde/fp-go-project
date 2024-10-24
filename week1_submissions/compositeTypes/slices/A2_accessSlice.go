package main

import "fmt"

func main() {
	// from Part 1
	var monthSpendings = []float64 {
		9.50, 8.00, 10.20, 7.40,
	}

	// Part 2
	fmt.Printf("Spent $%.2f in week 3\n", monthSpendings[2])

	monthSpendings[2] = 9.80
	fmt.Printf("Updated slice: %v\n", monthSpendings)
}