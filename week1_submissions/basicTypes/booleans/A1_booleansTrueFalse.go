package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var guess int
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&guess)
	
	var answer int = rand.Intn(100)
	if guess == answer {
		fmt.Println("Well Done! Your guess is correct")
	} else if guess < answer {
		fmt.Println("Too low, try again next time!")
	} else {
		fmt.Println("Too high, try again next time!")
	}
}
