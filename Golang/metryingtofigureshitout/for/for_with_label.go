package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var myStrings = `one
two
three
four
five`

	scanner := bufio.NewScanner(strings.NewReader(myStrings))
outerLoop:
	for scanner.Scan() {
		for _, char := range scanner.Text() {
			fmt.Println(string(char))
			if string(char) == "w" {
				fmt.Println("ITS A DABWIA!")
				continue outerLoop
			}
		}
		fmt.Println("Outer loop loops")

	}
}
