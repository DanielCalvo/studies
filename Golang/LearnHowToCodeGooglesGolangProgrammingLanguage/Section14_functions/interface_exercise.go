package main

import "fmt"

type cat struct {
	name       string
	age        int
	isfriendly bool
	likestuna  bool
}
type dog struct {
	name       string
	age        int
	isfriendly bool
	likesmango bool
}

//Anybody who has the function cuddle is also of type pet
type animal interface {
	pet()
}

func (c cat) pet() {
	fmt.Println("You've petted", c.name, "and it is very happy about it")
}
func (d dog) pet() {
	fmt.Println("You've petted", d.name, "and it is very happy about it")
}
func barbar(a animal) {

	switch a.(type) {
	//this is asserting: I'm asserting that this is of type pet and I checked it with the switch:
	case cat:
		fmt.Println("You've passed a cat named", a.(cat).name)
	case dog:
		fmt.Println("You've passed a dog named", a.(dog).name)
	}
}

func main() {

	jade := cat{
		name:       "Jade",
		age:        3,
		isfriendly: true,
		likestuna:  true,
	}
	shanti := dog{
		name: "Shanti",
	}

	fmt.Println(jade.name)
	jade.pet()
	shanti.pet()
	barbar(shanti)
	barbar(jade)
}
