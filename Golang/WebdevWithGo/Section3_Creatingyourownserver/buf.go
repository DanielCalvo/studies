package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "Here are\nSome lines\nSeparated\nBy newlines!"

	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
