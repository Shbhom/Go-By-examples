package main

import (
	"fmt"
)

func Main6() {

	var a [5]int
	fmt.Println("emp:", a) //every thing is initialized to 0
	a[1] = 2
	a[3] = 3
	i := 0
	for i < len(a) {
		if i != 0 {
			a[i] += a[i-1]
		}
		i++
	}
	fmt.Println(a)
	fmt.Println("lenght of a", len(a))

	b := [5]int{1, 2, 3, 5}
	fmt.Println(b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = a[i] + b[j]
		}
	}
	fmt.Println(twoD)
}
