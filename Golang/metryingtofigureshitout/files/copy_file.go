package main

import (
	"io"
	"os"
)

func main() {

	sourceFile, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create("/tmp/passwd")
	if err != nil {
		panic(err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		panic(err)
	}

}
