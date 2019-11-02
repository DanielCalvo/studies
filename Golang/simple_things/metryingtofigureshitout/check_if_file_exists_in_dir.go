package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	srcDir := "/tmp/img/"
	myFile := "IMG_20190826_223242.jpg"

	if findFileInDir(myFile, srcDir) {
		fmt.Println("Found!")
	} else {
		fmt.Println("not found :(")
	}

	//files, err := ioutil.ReadDir(srcDir)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, file := range files {
	//	if myFile == file.Name() {
	//		fmt.Println("Found myFile!", myFile)
	//	}
	//}

}

func findFileInDir(file string, dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range files {
		if file == v.Name() {
			return true
		}
	}
	return false
}
