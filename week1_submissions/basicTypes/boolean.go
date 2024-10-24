package main

import "fmt"

func main() {
	conditionA := 10
	conditionB := 12
	conditionC := 10
	if conditionA < conditionB {
		fmt.Println("A lesser than B")
	}
	if conditionB <= conditionC {
		fmt.Println("A lesser than C")
	}
	if conditionA == conditionC {
		fmt.Println("A is same as C")
	}
}
