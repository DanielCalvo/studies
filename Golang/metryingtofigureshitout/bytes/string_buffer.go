package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {

	s :=
		`Hey here's a multiline string
That I want to add some lines to!`

	scanner := bufio.NewScanner(strings.NewReader(s))

	var b bytes.Buffer

	for scanner.Scan() {
		scanner.Text()
		//You can do something with the string here!
		b.WriteString(scanner.Text() + "\n")
	}

	b.WriteString("Hello world!")

	fmt.Print(b.String())

}
