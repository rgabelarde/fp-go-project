package main

import "fmt"

func main() {
	var monthSpendings = []float64 {
		9.50, 8.00, 10.20, 7.40,
	}
	fmt.Printf("Length = %v, Capacity = %v", len(monthSpendings), cap(monthSpendings))
}