package main

import (
	"fmt"
	"math"
)

// interfaces are named collections of method signatures
// Whatever has the area function implements the geometry interface -- something like that
// Cool blog post on interfaces for when you want to learn more: https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
type geometry interface {
	area() float64
}

type rrectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rrectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) float64 {
	return g.area()
}

func main() {
	r := rrectangle{
		width:  4,
		height: 5,
	}

	fmt.Println(measure(r))
}
