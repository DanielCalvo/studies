package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
)

//add basic error checking

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Fatal("Error on JSON marshall:", err)
	}
	fmt.Println(string(bs))

}
