package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inputFile, err := os.Open("/tmp/nums/11.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())

}
