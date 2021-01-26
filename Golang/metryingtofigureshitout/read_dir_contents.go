package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	files, err := ioutil.ReadDir("/tmp/nums")
	if err != nil {
		panic(err)
	}
	fmt.Println("before")
	for _, v := range files {
		fmt.Println(v.Name())
	}
	fmt.Println("after")
}
