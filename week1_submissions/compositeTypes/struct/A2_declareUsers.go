package main

import "fmt"

type Customer struct {
	fname string
	lname string
	Age int
	Subscriber bool
	HomeAddress string
	Phone int
	CreditAvailable float32
	CurrentCardCost float32
	CurrentOrderCost float32
}

func main() {
	customer1 := Customer {
		fname: "Annakin",
		lname: "Skywalker",
		Age: 45,
		Subscriber: true,
		HomeAddress: "Death Star",
		Phone: 1234567,
		CreditAvailable: 10000.00,
		CurrentCardCost: 0,
		CurrentOrderCost: 0,
	}

	customer2 := Customer {
		fname: "Han",
		lname: "Solo",
		Age: 50,
		Subscriber: false,
		HomeAddress: "Tatooine",
		Phone: 4321765,
		CreditAvailable: 20000.00,
		CurrentCardCost: 0,
		CurrentOrderCost: 0,
	}

	fmt.Printf("Customer 1 Details:\n First Name: %v\n Last Name: %v\n Age: %v\n Subscriber: %v\n Home Address: %v\n Phone: %v\n Credit Available: %.2f\n Current Card Cost: %.2f\n Current Order Cost: %.2f\n\n", customer1.fname, customer1.lname, customer1.Age, customer1.Subscriber, customer1.HomeAddress, customer1.Phone, customer1.CreditAvailable, customer1.CurrentCardCost, customer1.CurrentOrderCost)
	fmt.Printf("Customer 2 Details:\n First Name: %v\n Last Name: %v\n Age: %v\n Subscriber: %v\n Home Address: %v\n Phone: %v\n Credit Available: %.2f\n Current Card Cost: %.2f\n Current Order Cost: %.2f\n", customer2.fname, customer2.lname, customer2.Age, customer2.Subscriber, customer2.HomeAddress, customer2.Phone, customer2.CreditAvailable, customer2.CurrentCardCost, customer2.CurrentOrderCost)
}