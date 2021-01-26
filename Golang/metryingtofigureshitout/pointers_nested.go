package main

import "fmt"

type Fruit struct {
	Name     string
	Weight   int
	IsTasty  bool
	Minerals Minerals
}

type Minerals struct {
	Potassium int
	Sodium    int
	Calcium   int
}

func ChangeFruit(f *Fruit) {
	f.Name = "Changed"
}

func ChangeMineral(f *Fruit) {
	f.Minerals.Calcium = 99
}

func main() {

	a := Fruit{
		Name:    "Banana",
		Weight:  2,
		IsTasty: true,
	}

	ChangeFruit(&a)
	ChangeMineral(&a)

	fmt.Println(a)
}
