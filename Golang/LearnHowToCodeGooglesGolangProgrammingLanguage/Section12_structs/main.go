package main

import (
	"fmt"
)

func main() {

	type person struct {
		first string
		last  string
		age   int
	}
	type secretAgent struct {
		person
		ltk bool
	}

	sa1 := secretAgent{
		person: person{
			first: "Dani",
			last:  "VeryDani",
			age:   25,
		},
		ltk: true,
	}
	//p2 := person{
	//	first : "Somefirstname",
	//	last: "Somesecondname",
	//}
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	//fmt.Println(p1, p2)

	//Attempt at anonymous struct
	a1 := struct {
		first, last string
		age         int
	}{
		first: "James",
		last:  "Bond",
		age:   40,
	}
	fmt.Println(a1)

}
