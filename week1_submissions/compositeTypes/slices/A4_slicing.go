package main

import "fmt"

func main() {
	// from Part 3
	var monthSpendings = []float64 {
		9.50, 8.00, 10.20, 7.40,
	}
	monthSpendings[2] = 9.80
	monthSpendings = append(monthSpendings, 8.40, 9.40, 7.20)

	// Part 4
	subSlice := monthSpendings[3:]
	fmt.Printf("Length = %v, Capacity = %v\n", len(subSlice), cap(subSlice))
}