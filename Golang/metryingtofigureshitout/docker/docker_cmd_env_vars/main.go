package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("name:", os.Getenv("name"))
	fmt.Println("surname:", os.Getenv("surname"))
	fmt.Println("age:", os.Getenv("age"))

}
