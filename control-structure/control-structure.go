package main

import (
	"fmt"
	"time"
)

func main() {
	a := 10

	// If statement
	if a > 0 {
		fmt.Println((a), "is a positive number")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is", i)
	}

	// switch statement
	decision := time.Now().Weekday().String()

	switch decision {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	case "Thursday":
		fmt.Println("Today is Thursday")
	case "Friday":
		fmt.Println("Today is Friday")
	case "Saturday":
		fmt.Println("Today is Saturday")
	case "Sunday":
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Invalid day")
	}

}
