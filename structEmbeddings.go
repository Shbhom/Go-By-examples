package main

import "fmt"

/*
kind of inheritance in Golang
so we have swe struct and description which embeds SWE

*/

type SWE struct {
	name string
	age  int
}

type description struct {
	SWE
	str string
}

func (s SWE) describe() string {
	return fmt.Sprintf("learing Golang")
}

type sweMeths interface {
	describe() string
}

func Main17() {
	// SWE1 := SWE{"shubhom", 20}
	// SWE1_desc := description{SWE1, "backend Developer learing Golang"}
	SWE1_desc := description{SWE{"shubhom", 20}, "backend Developer learing Golang"}

	var meth1 sweMeths = SWE1_desc
	//datatype of type sweMeth interface is also
	fmt.Println(meth1)

	fmt.Printf("developerName: %s, currentlyWorkingOn:%s", SWE1_desc.name, SWE1_desc.str)

}
