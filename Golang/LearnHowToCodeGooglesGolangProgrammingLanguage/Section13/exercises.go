package main

import "fmt"

func main() {

	//Exercise 1
	fmt.Println("Exercise 1")
	type person struct {
		firstname    string
		lastname     string
		fav_icecream []string
	}

	p1 := person{
		firstname:    "Dani",
		lastname:     "Verydani",
		fav_icecream: []string{"Chocolate", "Dark chocolate", "Cookie"},
	}
	p2 := person{
		firstname:    "Jade",
		lastname:     "Veryjade",
		fav_icecream: []string{"Cats don't eat icecream", "But they like meat"},
	}
	fmt.Println(p1)

	fmt.Println(p1.firstname, p1.lastname)

	for i := 0; i < len(p1.fav_icecream); i++ {
		fmt.Println(p1.fav_icecream[i])
	}

	//Exercise 2
	fmt.Println("Exercise 2")
	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	//Exercise 3
	fmt.Println("Exercise 3")

	type car struct {
		doors int
		color string
	}
	type truck struct {
		car
		fourWheel bool
	}
	type sedan struct {
		car
		luxury bool
	}

	danitruck := truck{
		car: car{
			doors: 2,
			color: "grey",
		},
		fourWheel: true,
	}
	fmt.Println(danitruck.doors, danitruck.color, danitruck.fourWheel)

	//Exercise 4
	//Extra challenge as sugested by Todd:
	//Use anonymous struct
	//Store one field being map, and another field being a slice
	fmt.Println("Exercise 4")

	a1 := struct {
		firstname  string
		lastname   string
		age        int
		pokemons   map[int]string
		chocolates []string
	}{
		firstname: "Dani",
		lastname:  "Verydani",
		age:       20,
		pokemons: map[int]string{
			2: "pikachu",
			3: "bulbasaur",
		},
		chocolates: []string{"Chocolate", "Cookies", "White chocolate"},
	}

	fmt.Println(a1)
}
