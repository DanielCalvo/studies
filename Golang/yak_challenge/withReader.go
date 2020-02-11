package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

//https://stackoverflow.com/questions/34654514/how-to-read-a-file-starting-from-a-specific-line-number-using-scanner

//read first line from file1
//read first line from file2
//if first line from file1 is smaller than from file2, move the scanner forward on file2, and make the scanner stay where it is on file 2.

func main() {

	inputFile, err := os.Open("/tmp/nums/11.txt")
	r := bufio.NewReader(inputFile)

	if err != nil {
		panic(err)
	}
	//var line []byte
	fPos := 0 // or saved position

	for i := 1; ; i++ {
		line, err := r.ReadBytes('\n')

		fmt.Printf("[line:%d pos:%d] %q\n", i, fPos, line)

		if err != nil {
			break
		}
		fPos += len(line)
	}

	if err != io.EOF {
		log.Fatal(err)
	}

}
