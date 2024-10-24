package main

import "fmt"

func main() {
	var measurements = [24]float64{
		20.1, 24, 27.3, 30.1, 26.4, 22.2, 20.1, 24, 27.3, 30.1,
		26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1,
		26.4, 20.1, 24, 27.3,
	}
	var sum, avg float64
	
	for i := 0; i < len(measurements); i++ {
		sum += measurements[i]
	}
	avg = sum / float64(len(measurements))
	fmt.Printf("Average temperature: %.2f\n", avg)
}