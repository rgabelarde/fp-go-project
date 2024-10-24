package main

import "fmt"

func main() {
	num1, num2 := 17, 24
	var bigger, difference int

	if num1 > num2 {
		bigger = num1
		difference = num1 - num2
	} else {
		bigger = num2
		difference = num2 - num1
	}

	fmt.Printf("The bigger number is: %d\n", bigger)
	fmt.Printf("It is bigger by: %d\n", difference)
}