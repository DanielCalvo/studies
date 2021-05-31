package main

import (
	"fmt"
	"github.com/google/go-github/github"
)

func main() {
	fmt.Println("https://github.com/google/go-github")
	client := github.NewClient(nil)
	fmt.Println(client)

}
