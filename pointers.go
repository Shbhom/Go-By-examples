package main

import (
	"fmt"
)

// type person struct {
// 	name string
// 	age  int
// }

// func makePerson() *person {
// 	m := person{"shubhom", 20}
// 	fmt.Printf("%p\n", &m)
// 	return &m
// }

func Main13() {
	//do display the use of heap in mem allocation
	// fmt.Printf("%p\n", makePerson())
	x, y := 1, 2
	//here the address is not changing because we are printing the addresses of x and y and not the actual address
	fmt.Printf("before swap %d,%d,%p,%p \n", x, y, &x, &y)
	swap(&x, &y)
	fmt.Printf("after swap %d,%d,%p,%p \n", x, y, &x, &y)
	// x, y = swapWithoutPointer(x, y)
	// fmt.Printf("again swapped without pointers %d,%d,%p,%p \n", x, y, &x, &y)
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func swapWithoutPointer(x, y int) (int, int) {
	temp := x
	x = y
	y = temp
	return x, y
}
