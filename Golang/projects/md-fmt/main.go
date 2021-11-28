package main

import (
	"bufio"
	"fmt"

	"os"
)

func main(){

	//Hey first use case! Kind of a dumb one!
	//open file
	//if a line starts with a minus sign, unindent it!
	//save the result to the file!

	//I wonder how go fmt works?

	//If line begins with whitespace!

	filepath := "/home/daniel/Projects/studies/Golang/projects/md-fmt/dummy.md"

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}


	}


}