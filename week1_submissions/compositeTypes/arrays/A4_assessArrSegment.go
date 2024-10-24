package main

import "fmt"

func main() {
	// Above from Part 3
	var arr = [6]string{
		"Windows XP",
		"Linux 1.0",
		"Raspbian Teensy",
		"MacOS",
		"IOS",
		"Google Android",
	}

	arr[0] = "Windows 10"
	arr[1] = "Linux 16.04"
	arr[2] = "Raspbian Buster"


	var s[]string = arr[0:]
	s = append(s, "Ubuntu", "MS-Dos", "Solaris")

	fmt.Println("Printing first 3 elements of the array %v", s[:3])
	fmt.Println("Printing next 3 elements of the array %v", s[3:6])
	fmt.Println("Printing last 3 elements of the array %v", s[len(s)-3:])

}