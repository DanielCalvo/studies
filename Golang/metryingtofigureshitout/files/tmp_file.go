package main

import (
	"fmt"
	"io/ioutil"
)

const tmpDir = "/tmp"

func main() {

	//Interesting, Go does not delete the file when the program finishes
	tempFile, err := ioutil.TempFile("", "myfile.txt")

	if err != nil {
		panic(err)
	}

	_, err = tempFile.WriteString("I am writing to my temporary file!")
	fmt.Println("Wrote to temporary file: ", tempFile.Name())
	if err != nil {
		panic(err)
	}
}
