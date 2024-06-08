package main

import (
	"fmt"
)

type Rectangle struct {
	Width, Height int
}

// In Go, type definitions and method definitions should be at the package level,
// not inside another function.
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	fmt.Println("Go Topics")
	fmt.Println("example of arrays and slices")

	var arr [3]int = [3]int{1, 2, 3}

	fmt.Println("array arr: ", arr)

	fmt.Println("example of slices")

	slice := append(arr[:], 4)

	fmt.Println("slice: ", slice)

	fmt.Println("example of maps")

	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2

	fmt.Println("map m: ", m)

	fmt.Println("example of structs")

	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}

	fmt.Println("struct p: ", p)

	fmt.Println("example of methods and interfaces")

	r1 := Rectangle{Width: 40, Height: 20}
	fmt.Println("Area of rectangle r1: ", r1.Area())
}
