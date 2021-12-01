package main

import (
	"bufio"
	"fmt"
	"strings"

	"os"
)

func main() {
	filepath := "/home/daniel/Projects/studies/Golang/projects/md-fmt/dummy.md"

	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var line string

	for scanner.Scan() {
		line = scanner.Text()

		if strings.HasPrefix(line, "###") {
			fmt.Println(line, "<- So this is the line I want")
			//I think I have to manipulate line here!
			//I wonder how gofmt did it?
		} else {
			fmt.Println(line)
		}

	}

}
