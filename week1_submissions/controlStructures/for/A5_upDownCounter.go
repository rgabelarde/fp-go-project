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

	fmt.Println("Counting up: ")
	for i := startVal; i <= endVal; i++ {
		fmt.Println(i)
	}

	fmt.Println("Counting down: ")
	for i := endVal; i >= startVal; i-- {
		fmt.Println(i)
	}
}