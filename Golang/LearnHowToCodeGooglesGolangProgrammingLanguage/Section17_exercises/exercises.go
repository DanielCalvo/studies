package main

import "fmt"

//Exercise 2
type person struct {
	first string
	last  string
	age   int
}

func main() {

	//Exercise 1
	a := "value"
	fmt.Println(&a)

	//Exercise 2
	p1 := person{
		first: "Dani",
	}
	fmt.Println(p1.first)
	changeMe(&p1)
	fmt.Println(p1.first)
}

func changeMe(p *person) {
	p.first = "Danidani"
	//(*p).first = "Danidanidani" //Also valid!
}
