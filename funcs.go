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

// variadic functions
func variadicFunc(nums ...int) {
	fmt.Println("len of variadic arguments", len(nums))
	for _, ele := range nums {
		fmt.Print(ele)
	}
	fmt.Println()
}

func Main10() {
	fmt.Println(mult(1, 2))
	fmt.Println(add(1, 2))
	fmt.Println("multiValued Return")
	value, Ty := multiReturn(1, 2)
	fmt.Println(value, Ty)
	fmt.Println("variadic Function")
	variadicFunc(1, 2, 3, 4, 5)
}
