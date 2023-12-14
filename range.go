package main

import "fmt"

/*
range is a keyword ~= to enumurate in python
it helps in iterating over the data structure
*/

func Main9() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i, ele := range arr {
		fmt.Println("index and it's value :", i, "->", ele)
	}

	m := map[string]int{"first": 1, "sec": 2, "third": 3}

	for k, v := range m {
		fmt.Println("key and its value:", k, ":", v)
	}

	//now lets say we need only the values or only the key from the map

	for key := range m {
		fmt.Println(key)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	//if we have a string so we can iterate over each character and it give use its unicode points

	for _, c := range "hello world" {
		fmt.Print(c, "\n")
	}

}
