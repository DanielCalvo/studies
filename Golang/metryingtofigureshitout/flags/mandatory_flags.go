package main

import (
	"flag"
	"fmt"
)

func main() {
	//go run mandatory_flags.go -name=Bob

	nameFlag := flag.String("name", "", "a string")

	flag.Parse()
	fmt.Println("Name is:", *nameFlag)

	if *nameFlag == "" {
		panic("Name can't be empty!")
	}
	fmt.Println("Good morning", *nameFlag)

}
