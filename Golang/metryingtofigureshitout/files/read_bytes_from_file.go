package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}

	var line []byte

	//let's read things until we find a newline character
	for {
		b := make([]byte, 1)

		_, err = file.Read(b)
		if err != nil {
			panic(err)
		}

		if b[0] == '\n' {
			break
		}
		line = append(line, b[0])
	}
	fmt.Println(string(line))

}
