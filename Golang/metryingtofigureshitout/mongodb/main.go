package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func main() {
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
