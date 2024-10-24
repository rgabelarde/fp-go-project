package main

import "fmt"

func main() {
	var roomTemps = [][]float64{
		{20, 21, 23, 25, 22},
		{27, 23, 25, 20, 30, 24},
		{22, 23, 24, 22},
	}

	var roomAvgs []float64

	for i := 0; i < len(roomTemps); i++ {
		var sum float64
		for j := 0; j < len(roomTemps[i]); j++ {
			sum += roomTemps[i][j]
		}
		roomAvgs = append(roomAvgs, sum/float64(len(roomTemps[i])))
	}

	fmt.Printf("Room 1 average temperature: %.2f\n", roomAvgs[0])
	fmt.Printf("Room 2 average temperature: %.2f\n", roomAvgs[1])
	fmt.Printf("Room 3 average temperature: %.2f\n", roomAvgs[2])
}
