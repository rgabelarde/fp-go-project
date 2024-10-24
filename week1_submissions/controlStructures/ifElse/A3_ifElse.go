package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter an integer value: ")
	fmt.Scanln(&num)

	if num % 5 == 0 && num % 6 == 0 {
		fmt.Println("The number is divisible by both 5 and 6")
	} else {
		fmt.Println("The number is NOT divisible by 5 AND 6")
	}
}