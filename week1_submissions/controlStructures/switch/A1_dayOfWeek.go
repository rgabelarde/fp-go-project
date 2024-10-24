package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	switch today {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 0:
			fmt.Println("Sunday")
	}
	if today % 2 == 0 {
		fmt.Println("Today is an even day")
	} else {
		fmt.Println("Today is an odd day")
	}

}