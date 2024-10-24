package main

import (
	"fmt"
)

func main() {
	var formatNo int
	var temp float64

	fmt.Println("Which format would you like to input your temperature (1: Kelvin, 2: Celsius, 3: Fahrenheit): ")
	fmt.Scanln(&formatNo)

	fmt.Println("Enter the current temperature value: ")
	fmt.Scanln(&temp)

	var otherFormat1, otherFormat2 float64
	switch formatNo {
		case 1:
			otherFormat1 = temp - 273.15
			otherFormat2 = (temp - 273.15) * 9/5 + 32
			fmt.Println("Temperature in Celsius: ", otherFormat1)
			fmt.Println("Temperature in Fahrenheit: ", otherFormat2)

		case 2:
			otherFormat1 = temp + 273.15
			otherFormat2 = temp * 9/5 + 32
			fmt.Println("Temperature in Kelvin: ", otherFormat1)
			fmt.Println("Temperature in Fahrenheit: ", otherFormat2)

		case 3:
			otherFormat1 = (temp - 32) * 5/9 + 273.15
			otherFormat2 = (temp - 32) * 5/9
			fmt.Println("Temperature in Kelvin: ", otherFormat1)
			fmt.Println("Temperature in Celsius: ", otherFormat2)

		default:
			fmt.Println("Invalid format number")
			return
	}

}