package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {

	a := []string{
		"/tmp/mdscanner/kompose/CHANGELOG.md",
		"/tmp/mdscanner/kompose/something/123-minutes.md",
	}

	b := []string{
		"changelog",
		"minutes",
	}

	fmt.Println(strings.ToLower(a))
	fmt.Println(b)

	x := strings.Contains(strings.ToLower(a), b)

	if x {
		log.Println("yes")
	}
}
