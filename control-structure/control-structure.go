package main

import (
	"fmt"
)

func main() {
	a := 10

	if a > 0 {
		fmt.Println((a), "is a positive number")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is", i)
	}
}
