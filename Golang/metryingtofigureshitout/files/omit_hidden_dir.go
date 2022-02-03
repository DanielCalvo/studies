package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
mkdir /tmp/a
mkdir /tmp/a/banana
mkdir /tmp/a/.apple
touch /tmp/a/.orange
*/
func main() {

	folders, err := ioutil.ReadDir("/tmp/a")
	if err != nil {
		panic(err)
	}

	for _, folder := range folders {
		if !strings.HasPrefix(folder.Name(), ".") {
			fmt.Println(folder.Name())
		}
	}
}
