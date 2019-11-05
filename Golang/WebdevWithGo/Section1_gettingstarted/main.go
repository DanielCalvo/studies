package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type superHero struct {
	person
	Superpower string
}

func (p person) speak() {
	fmt.Println("Hello, I'm", p.Name)
}

func (s superHero) speak() {
	fmt.Println("Hello, I'm", s.Name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)

	x := map[string]int{
		"Dani":   00,
		"McDani": 01,
	}
	fmt.Println(x)

	p1 := person{
		"Dani",
		00,
	}
	fmt.Println(p1)

	p2 := person{
		Name: "Dani",
		Age:  00,
	}
	fmt.Println(p2)
	p2.speak()

	s1 := superHero{
		person{
			Name: "Dani",
		},
		"banana",
	}
	fmt.Println(s1)

	s2 := superHero{
		person{
			"Dani dani",
			00,
		},
		"",
	}
	s2.speak()
	saySomething(s2)

}

// func (receiver) identifier(parameters) (returns) {
//Code
//}
