package main

import "fmt"

/*
just like any OOP supporting languages we have classes which encapsulates the data and methods
In go we can createMethods on a structure which could modify and apply changes that struct for us
*/

type person struct {
	name string
	age  int32
}

func (p *person) setter(age int32, name string) *person {
	if len(name) == 0 {
		p.age = age
	} else if age == 0 {
		p.name = name
	} else {
		p.age = age
		p.name = name
	}
	return p
}

func makePerson(name string, age int32) *person {
	p := person{name, age}
	fmt.Printf("pointer inside the function call: %p\n", &p)
	return &p
}

/*
as the makePerson call ends the p is stored or the address  to p is stored in the heap memory
not similar to the  heap DS and as we are stores the returned  value of the fucn
*/

func Main15() {
	// to create a person object from the struct
	per1 := makePerson("shivom", 22)
	fmt.Printf("pointer inside main%p\n", per1)
	fmt.Println(person{"shubhom", 23})

	per1 = per1.setter(43, "al;skdjf")
	fmt.Println(per1)

	//an anonymous struct when we have to use a struct for a single use only we don't have to define it
	dog := struct {
		name   string
		isGood bool
	}{"swift", true}
	fmt.Println(dog)

}
