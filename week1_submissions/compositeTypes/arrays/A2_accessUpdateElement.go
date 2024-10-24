package main

import "fmt"

func main() {
	// from Part 1
	var arr = [6]string{
		"Windows XP",
		"Linux 1.0",
		"Raspbian Teensy",
		"MacOS",
		"IOS",
		"Google Android",
	}

	// Part 2
	arr[0] = "Windows 10"
	arr[1] = "Linux 16.04"
	arr[2] = "Raspbian Buster"

	fmt.Println("Updated Array:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}