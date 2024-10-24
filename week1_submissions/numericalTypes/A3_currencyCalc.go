package main

import (
	"fmt"
	"math"
)

func main() {
	var no1DollarCoins, no50CentCoins, no20CentCoins, no10CentCoins, no5CentCoins int
	var totalAmount float64

	fmt.Println("Enter the number of $1 coins: ")
	fmt.Scanln(&no1DollarCoins)

	fmt.Println("Enter the number of 50c coins: ")
	fmt.Scanln(&no50CentCoins)

	fmt.Println("Enter the number of 20c coins: ")
	fmt.Scanln(&no20CentCoins)

	fmt.Println("Enter the number of 10c coins: ")
	fmt.Scanln(&no10CentCoins)

	fmt.Println("Enter the number of 5c coins: ")
	fmt.Scanln(&no5CentCoins)

	totalAmount = float64(no1DollarCoins) + float64(no50CentCoins) * 0.5 + float64(no20CentCoins) * 0.2 + float64(no10CentCoins) * 0.1 + float64(no5CentCoins) * 0.05
	fmt.Printf("Total amount of money: %.2f\n", totalAmount)

	fmt.Printf("Number of $2 notes you can get in exchange: %v ($%.2f change)", int(math.Floor(totalAmount / 2)), math.Mod(totalAmount, 2))
}