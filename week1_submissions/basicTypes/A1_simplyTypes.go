package main

import "fmt"

func main() {
	accText := "The following is the account information"
	firstName := "Luke"
	lastName := "Skywalkter"
	age := 20
	weight := 73
	height := 1.72
	remainingCredits := 123.55
	accName := "admin"
	accPassword := "password"
	isSubscribedUser := true

	fmt.Printf("%v  (type: %T)\n", accText)
	fmt.Printf("First Name: %v (type: %T)\n", firstName, firstName)
	fmt.Printf("Last Name: %v (type: %T)\n", lastName, lastName)
	fmt.Printf("Age: %v years old (type: %T)\n", age, age)
	fmt.Printf("Weight: %vkg (type: %T)\n", weight, weight)
	fmt.Printf("Height: %vm (type: %T)\n", height, height)
	fmt.Printf("Remaining Credits: $%v (type: %T)\n", remainingCredits, remainingCredits)
	fmt.Printf("Account Name: %v (type: %T)\n", accName, accName)
	fmt.Printf("Account Password: %v (type: %T)\n", accPassword, accPassword)
	fmt.Printf("Is Subscribed User: %v (type: %T)\n", isSubscribedUser, isSubscribedUser)
}
