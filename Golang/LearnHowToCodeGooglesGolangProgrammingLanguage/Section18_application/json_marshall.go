package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "Dani",
		Last:  "Verydani",
		Age:   42,
	}

	p2 := person{
		First: "Daniclone",
		Last:  "Almostdani",
		Age:   42,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
