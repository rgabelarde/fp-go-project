package main

import (
	"fmt"
)

func main() {
	var startVal, endVal int
	
	fmt.Print("Enter a lower limit: ")
	fmt.Scanln(&startVal)

	fmt.Print("Enter an upper limit: ")
	fmt.Scanln(&endVal)

	fmt.Printf("Even numbers between %v and %v are: ", startVal, endVal)
	for i := startVal; i <= endVal; i++ {
		if i%2 == 0 {
			if i > startVal && i != startVal+1 {
				fmt.Print(", ")
			}
			fmt.Print(i)
		}
	}

	fmt.Printf("\nOdd numbers between %v and %v are: ", startVal, endVal)
	for i := startVal; i <= endVal; i++ {
		if i%2 != 0 {
			if i > startVal && i != startVal+1 {
				fmt.Print(", ")
			}
			fmt.Print(i)
		}
	}

}