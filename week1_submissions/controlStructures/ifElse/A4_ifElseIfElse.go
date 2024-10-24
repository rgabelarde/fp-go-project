package main

import (
	"fmt"
)

func main() {
	var nameInput string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&nameInput)

	if nameInput == "Admin" {
		fmt.Println("Welcome Admin!")
	} else if nameInput == "Robin" || nameInput == "John" {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("Merry Men")
	}
}