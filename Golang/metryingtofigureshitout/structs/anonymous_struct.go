package main

import "fmt"

func main() {
	people := []struct {
		name    string
		age     int
		address string
	}{
		{name: "John Doe", age: 30, address: "123 Main St"},
		{name: "Bob", age: 31, address: "Bob's Main St"},
		//For clarity it is best if you do it like it is above, but you can also omit the field names:
		{"Bob McBobson", 13, "Banana st"},
	}

	//Lets do just a person as opposed to a slice:
	person := struct {
		name    string
		age     int
		address string
	}{
		name:    "bob",
		age:     0,
		address: "bob's street",
	}

	fmt.Printf("%#v\n", people)
}
