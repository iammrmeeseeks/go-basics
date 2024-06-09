package main

import (
	"errors"
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

// interfaces
type Animal interface {
	Sound() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Sound() string {
	return "Woof"
}

func (c Cat) Sound() string {
	return "Meow"
}

func MakeSound(a Animal) string {
	return a.Sound()
}

// error handling
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
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

	fmt.Println("example of methods")

	r1 := Rectangle{Width: 40, Height: 20}
	fmt.Println("Area of rectangle r1: ", r1.Area())

	fmt.Println("example interfaces")

	d := Dog{Name: "Max"}
	c := Cat{Name: "Whiskers"}

	fmt.Println("Dog sound: ", MakeSound(d))
	fmt.Println("Cat sound: ", MakeSound(c))

	fmt.Println("example of pointers")
	var x int = 10
	var pointer *int = &x

	fmt.Println("value of x: ", x)
	fmt.Println("pointer pointer: ", pointer)
	fmt.Println("value at pointer po: ", *pointer)

	fmt.Println("example of error handling")

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("result: ", result)
	}
}
