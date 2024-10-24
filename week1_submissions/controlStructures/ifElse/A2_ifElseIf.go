package main

import (
	"fmt"
)

func main() {
	num := 17

	if num % 2 == 0 {
		fmt.Println("The number is even")
	} else if num % 2 != 0 {
		fmt.Println("The number is odd")
	} 

	if num / 10 > 0 {
		fmt.Println("The number has more than one digit")
	} else if num / 10 == 0 {
		fmt.Println("The number has only one digit")
	}
}