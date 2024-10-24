package main

import "fmt"

func main() {
	var arr = [6]string{
		"Windows XP",
		"Linux 1.0",
		"Raspbian Teensy",
		"MacOS",
		"IOS",
		"Google Android",
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v, length: %v\n", arr[i], len(arr[i]))
	}

}