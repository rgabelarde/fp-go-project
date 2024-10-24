package main

import (
	"fmt"
)

func main() {
	for {
		var number int
		fmt.Print("Input a number: ")
		fmt.Scanln(&number)

		if number % 2 == 0 {
			fmt.Println("The number is even")
		} else {
			fmt.Println("The number is odd")
		}

		if number / 10 != 0 {
			fmt.Println("The number has more than one digit")
		} else {
			fmt.Println("The number has only one digit")
		}
	}
}