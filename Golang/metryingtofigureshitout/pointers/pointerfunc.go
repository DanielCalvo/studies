package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	//Now hold up, this is how it works right?
	p := Person{
		name: "Bob",
		age:  1,
	}
	ChangeByPointer(&p)

	fmt.Println(p) //Yeah, changed by pointer instead of returning something

}

func ChangeByPointer(p *Person) {
	p.age = 42
}
