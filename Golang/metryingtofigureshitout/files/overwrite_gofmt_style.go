package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("yo")

	filename := "/home/daniel/Projects/myfile.txt"

	src, err := os.ReadFile(filename) //Seems to be a bit more abstracted than io.ReadAll
	if err != nil {
		panic(err)
	}

	//Let's add something to the src []byte:
	b := []byte("Hello world!\n")
	src = append(src, b...)

	err = os.WriteFile(filename, src, 0644)
	if err != nil {
		panic(err)
	}

}
