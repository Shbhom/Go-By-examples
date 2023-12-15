package main

import (
	"fmt"
	"math"
)

/*

providing go a more OOP approach

interfaces are collection of methods so if a type has all these methods implemented then
it can be used

for eg if in circ instead of peri we had name Perimeter then we cannot use it with the interface
geometry and if I have any struct or type which have funcs area and peri can be used with interface geometry
*/

type geometry interface {
	area() float64
	peri() float64
}

type rect struct {
	length, bredth float64
}
type circ struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.length * r.bredth
}
func (r *rect) peri() float64 {
	return 2 * (r.length + r.bredth)
}

func (r *circ) area() float64 {
	return math.Pi * r.radius * r.radius
}
func (r *circ) peri() float64 {
	return 2 * math.Pi * r.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.peri())
}

func Main16() {
	rect := rect{10.00, 12.12}
	circ := circ{5.43}
	measure(&rect)
	measure(&circ)
}
