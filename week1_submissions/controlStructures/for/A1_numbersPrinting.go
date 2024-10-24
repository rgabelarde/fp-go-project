package main

import (
	"fmt"
)

func main() {
	fmt.Println("Printing numbers from 0 to 1000")
	for i := 0; i <= 1000; i++ {
		fmt.Println(i)
	}

	fmt.Println("Printing even numbers from 0 to 1000")
	for i := 0; i <= 1000; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("Printing odd numbers from 0 to 1000")
	for i := 0; i <= 1000; i++ {
		if i % 2 != 0 {
			fmt.Println(i)
		}
	}
}