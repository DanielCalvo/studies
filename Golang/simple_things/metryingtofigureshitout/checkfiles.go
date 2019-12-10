package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		_, _ = os.Stat("/tmp/" + file.Name())
	}
}
