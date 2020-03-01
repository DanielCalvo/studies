package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

//https://stackoverflow.com/questions/34654514/how-to-read-a-file-starting-from-a-specific-line-number-using-scanner

//read first line from file1
//read first line from file2
//if first line from file1 is smaller than from file2, move the scanner forward on file2, and make the scanner stay where it is on file 2.

//make an object with a function: GetnextLine and run it whenever you have to?

//type Mylist struct {
//	filePath string
//	name string
//	currentValue int
//}
//
//func (m Mylist) GetNextValue(n int) {
//	m.currentValue = n
//}
//
//func (m Mylist) New(filePath string) Mylist {
//
//}

func main() {

	inputFile, err := os.Open("/tmp/nums/11.txt")
	r := bufio.NewReader(inputFile)

	if err != nil {
		panic(err)
	}
	//var line []byte
	fPos := 0 // or saved position

	line, _, err := r.ReadLine()

	aByteToInt, _ := strconv.Atoi(string(line))
	fmt.Println(aByteToInt)

	if err != nil {
		panic(err)
	}
	fPos += len(line)
	inputFile.Seek(0, 0)

	if err != io.EOF {
		log.Fatal(err)
	}

}
