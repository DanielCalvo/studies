package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Iterate over a file and hold the following in a struct:
- The previous line
- The current line
- The next line
*/

type Line struct {
	previous string
	current  string
	next     string
}

func main() {
	src, err := os.ReadFile("/etc/passwd")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(src)))

	var line Line
	counter := -1
	for scanner.Scan() {
		counter++
		line.previous = line.current
		line.current = line.next
		line.next = scanner.Text()

		if line.current != "" {
			fmt.Println(counter, "-", line.current)
		}
	}

	//The scanner has reached its end above, but we must still process the last line
	counter++
	line.previous = line.current
	line.current = line.next
	line.next = ""
	fmt.Println(counter, "-", line.current)

}
