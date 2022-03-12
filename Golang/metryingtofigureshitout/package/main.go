package main

/*
go mod init example
go mod tidy
*/

import (
	"example/banana"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Banana")
	log.Println("also banana")
	banana.Print("Aha!")
}
