package main

import (
	"fmt"
	"reflect"
)

func mult(x int, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

//multi-value return function

func multiReturn(x, y int) (int, string) {
	Type := reflect.TypeOf(x).String()
	return x * y, Type
}

func main() {
	fmt.Println(mult(1, 2))
	fmt.Println(add(1, 2))
	fmt.Println(multiReturn(1, 2))

	value, _ := multiReturn(1, 2)
	fmt.Println(value)
}
