package main

import (
	"fmt"
)

func main() {
	var height, weight float64

	fmt.Print("Enter your height in meters: ")
	fmt.Scanln(&height)

	fmt.Print("Enter your weight in kilograms: ")
	fmt.Scanln(&weight)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI is %.2f\nBased on your BMI, you are ", bmi)

	switch {
	case bmi < 18.5:
		fmt.Println("underweight")
	case bmi >= 18.5 && bmi < 25:
		fmt.Println("healthy weight")
	case bmi >= 25 && bmi < 30:
		fmt.Println("overweight")
	case bmi >= 30 && bmi < 35:
		fmt.Println("Obese")
	case bmi >= 35 && bmi < 40:
		fmt.Println("severely obese")
	default:
		fmt.Println("morbidly obese")
	}
}
