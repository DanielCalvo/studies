package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	md := `### Hello!
You have opened dummy.md. I hope you have a good day!
- one item
 - one item that is not well indented :o`

	//strings.NewReader, neat!
	scanner := bufio.NewScanner(strings.NewReader(md))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
