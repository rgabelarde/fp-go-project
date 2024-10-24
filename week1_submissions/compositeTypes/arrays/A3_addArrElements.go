package main

import "fmt"

func main() {
	// from Part 2
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

	// part 3
	var s[]string = arr[0:]
	s = append(s, "Ubuntu", "MS-Dos", "Solaris")

	fmt.Println("Updated Array:")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}