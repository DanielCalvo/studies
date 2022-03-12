package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	//Looks like you need to load the json from the file first and unmarshall it into person!
	var p []Person

	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println("Unable to open file! I'm guessing it doesn't exist, let's continue! You could handle this better")
	}

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Unable to readall!")
	}

	err = json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("Unable to unmarshall! Could it be that the file does not exist?")
	}

	p1 := Person{
		Name: "Bob",
		Age:  10,
	}

	p2 := Person{
		Name: "Joe McJoeson",
		Age:  40,
	}

	p = append(p, p1)
	p = append(p, p2)

	fmt.Println(p)

	b, err = json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	//Tremendously crude, overwrites the entire file every time we just want to add a json element but... oh well
	err = os.WriteFile("test.json", b, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

}
