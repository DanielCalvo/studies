package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filepath := "/home/daniel/Projects/studies/Golang/projects/md-fmt/dummy.md"

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		//we need some regex in here, this doesn't cut it
		if strings.HasPrefix(line, " ") {
			fmt.Println(line)
		}
	}

}
