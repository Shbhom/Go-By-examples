package main

import "fmt"

func makeCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func Main11() {
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Println("counter 1 current value", counter1())
	fmt.Println("counter 1 current value", counter1())
	fmt.Println("counter 1 current value", counter1())
	fmt.Println("counter 2 current value", counter2())
	fmt.Println("counter 2 current value", counter2())
}
