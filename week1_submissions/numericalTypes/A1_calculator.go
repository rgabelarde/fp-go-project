package main

import "fmt"

func main() {
	var input1, input2 int
	var arithmaticFunction string

	fmt.Println("Enter an integer between 1 and 9: ")
	fmt.Scanln(&input1)

	fmt.Println("Enter an arithmatic function of choice (+, -, /, *): ")
	fmt.Scanln(&arithmaticFunction)

	fmt.Println("Enter another integer between 1 and 9: ")
	fmt.Scanln(&input2)

	var result float64
	switch arithmaticFunction {
		case "+":
			result = float64(input1 + input2)
		case "-":
			result = float64(input1 - input2)
		case "/":
			result = float64(input1 / input2)
		case "*":
			result = float64(input1 * input2)
		default:
			fmt.Println("Invalid arithmatic function")
			return
	}
	fmt.Printf("Your result is: %.1f", result)

}