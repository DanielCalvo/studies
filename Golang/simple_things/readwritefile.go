package main

import (
	"fmt"
	"io/ioutil"
)

//func check(e error) {
//	if e != nil {
//		panic(e)
//	}
//}

func main() {
	filelocation := "/home/daniel/PycharmProjects/studies/Golang/simple_things/myfile.txt"
	//Printing a text file
	dat, err := ioutil.ReadFile(filelocation)
	check(err)
	fmt.Println(string(dat))

	//Simple write to file
	d1 := []byte("hello\ngo\n") //To begin with, first thing we need is a slice of bytes
	err = ioutil.WriteFile(filelocation, d1, 0644)
	check(err)

	//How to append to a file again?
}
