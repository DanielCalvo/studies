package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Animal struct {
	Species string
	Age     int
}

func main() {

	frank := Animal{
		Species: "dog",
		Age:     1,
	}

	mike := Animal{
		Species: "cat",
		Age:     3,
	}

	var animals []Animal

	animals = append(animals, frank)
	animals = append(animals, mike)

	fmt.Println(frank, mike)
	fmt.Println(animals)

	myJson, err := json.MarshalIndent(animals, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/tmp/myJson.json", myJson, 0644)

}
