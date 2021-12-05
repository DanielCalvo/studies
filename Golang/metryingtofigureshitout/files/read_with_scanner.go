package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//The idiomatic way to read lines from a file in Go!
	file, err := os.Open("/tmp/a.txt")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
