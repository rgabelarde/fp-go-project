package main

import (
	"fmt"
)

func main() {
	var numInput int

	for {
		fmt.Print("Enter a number: ")
		fmt.Scanln(&numInput)

		if numInput % 2 == 0 {
			fmt.Printf("The number %v is even\n", numInput)
		} else {
			fmt.Printf("The number %v is odd\n", numInput)
		}
	}
}