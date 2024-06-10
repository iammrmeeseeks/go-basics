package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	c := make(chan int)

	go sum(a, c)
	fmt.Println(<-c)
}
