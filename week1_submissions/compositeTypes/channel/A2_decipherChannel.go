// Output for Activity 2:
// 1
// 2

// Overfilling input channel buffer
package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with a capacity of 2
	ch := make(chan int, 2)

	// Send two values to fill the channel
	ch <- 1
	ch <- 2

	// This third send will cause a panic due to overflow
	ch <- 3 // Attempting to overflow the channel

	// This line won't be reached due to the panic above
	fmt.Println("This won't be printed.")
}
