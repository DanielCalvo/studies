package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//If you run without arguments, the user will default to daniel (your user)
	//But you can also: go run main.go -name=Bob
	nameFlag := flag.String("name", os.Getenv("USER"), "a string")
	flag.Parse()

	fmt.Println("Nameflag:", *nameFlag)
}
