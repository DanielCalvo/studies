package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speaking")
}
func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

type human interface {
	speak()
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier (parameters) (return) { code }
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

//Keyword, identifier, type

func main() {
	foo()
	s1 := woo("Apple")
	fmt.Println(s1)
	x, y := mouse("Dani", "Verydani")
	fmt.Println(x, y)
	a := sum(1, 2, 3, 4, 5)
	fmt.Println(a)
	xi := []int{1, 2, 3, 4, 5}
	sum(xi...)
	fmt.Println()
	//defer foo()
	aa := sum(2)
	fmt.Println(aa)

	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
	}
	fmt.Println(sa1)
	sa1.speak()

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(p1)

	bar(sa1)
	bar(p1)

	//conversion
	//type hotdog
	//var bbb hotdog = 42
	//aaa := int(bbb)
	//fmt.Printf("%T", aaa)

}

//func (r receiver) identifier (parameters) (return) { ... }

func foo() {
	fmt.Println("Hello from foo")
}

func woo(s string) string {
	return s
}

func mouse(fn string, ln string) (string, bool) {
	fmt.Println("Hello from mouse!", fn, ln)
	return fn, true
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
	return sum
}
