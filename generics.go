package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
helps in maintaing DRY principal in our code
i.e. don't repeat yourself
*/

// func add(x, y int) int {
// 	return x + y
// }

func addWithGenerics[T constraints.Ordered](x, y T) T {
	return x + y
}

type customMap[T comparable, v any] map[T]any

//defining a linked List and its' node

type node[T any] struct {
	next *node[T]
	val  T
}

type list[T any] struct {
	head, tail *node[T]
}

//a function to add new elements to our LL

func (l *list[T]) push(v T) {
	if l.tail == nil {
		l.head = &node[T]{val: v}
		l.tail = l.head
	} else {
		l.tail.next = &node[T]{val: v}
		l.tail = l.tail.next
	}
}

func (l *list[T]) printLL() []T {
	var arr []T
	for e := l.head; e != nil; e = e.next {
		arr = append(arr, e.val)
	}
	return arr
}

//here this comparable is any datatype which can be compared i.e int==int || float==float

func main() {
	fmt.Println(addWithGenerics[int](1, 2))
	fmt.Println(addWithGenerics[float64](1.1231, 2.1231))
	fmt.Println(addWithGenerics("adskfj", "asdfkjas"))
	//the aboveFunction won't work with floats or any other int type also
	// var a, b int32 = 4, 5
	// fmt.Println(add(a, b))

	Map1 := customMap[int, string]{1: "one"}
	Map2 := customMap[string, int]{"one": 1}
	fmt.Println(Map1, Map2)

	list1 := list[int]{}

	fmt.Println(list1.printLL())
	list1.push(23)
	list1.push(12)
	fmt.Println(list1.printLL())

	list2 := list[string]{}

	fmt.Println(list2.printLL())
	list2.push("sldfk")
	list2.push("la")
	list2.push(";sdkfk")
	fmt.Println(list2.printLL())

}
