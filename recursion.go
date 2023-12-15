package main

import "fmt"

func Facto(x int) int {
	if x == 0 || x == 1 {
		return 1
	} else {
		return x * Facto(x-1)
	}
}

// return nth fibo term
func fibo(x int) int {
	if x < 2 {
		return x
	} else {
		return fibo(x-2) + fibo(x-1)
	}
}

func Main12() {
	fmt.Println(Facto(5))
	fmt.Println(fibo(3))
}
