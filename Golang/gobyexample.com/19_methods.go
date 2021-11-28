package main

import "fmt"

//Go supports methods defined on struct types

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func main() {

	r := rectangle{2, 3}

	//Go automatically handles conversion between values and pointers for method calls
	fmt.Println(r.area())

}
