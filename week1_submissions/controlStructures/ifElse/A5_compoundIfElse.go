package main

import (
	"fmt"
)

func main() {
	var year int

	fmt.Print("Enter a year: ")
	fmt.Scanln(&year)

	if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a Leap Year")
	}

}