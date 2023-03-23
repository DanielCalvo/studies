package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p //A pointer to a local variable will survive the scope of the function
}

// It’s idiomatic to encapsulate new struct creation in constructor functions
func newPersonButWithNoPointer(name string) person {
	p := person{name: name}
	p.age = 42
	return p
}

func main() {
	fmt.Println(person{name: "Alice", age: 42})

	p1 := newPerson("Someone")
	fmt.Println(*p1)
	p2 := newPersonButWithNoPointer("peeeeerson")
	fmt.Println(p2)
}
