package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var guess int
	
	answer := rand.Intn(100)

	for i := 1; i <= 5; i++ {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		
		if guess == answer {
			fmt.Print("Well Done! Your guess is correct")
			break
		} else if guess == 101 {
			fmt.Print("You have given up! The answer is", answer)
			break
		} else if i == 5 {
			fmt.Print("Wrong answer :( You have exhausted all your tries! The answer is: ", answer)
			break
		} else if guess < answer {
			fmt.Println("Too low, try again! You have", 5-i, "tries left")
		} else if guess > answer{
			fmt.Println("Too high, try again! You have", 5-i, "tries left")
		}
	}
}