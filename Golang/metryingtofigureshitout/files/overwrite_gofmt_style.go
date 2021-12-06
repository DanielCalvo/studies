package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("yo")

	myFile, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}

	src, err := io.ReadAll(myFile)
	fmt.Print(string(src))
}
