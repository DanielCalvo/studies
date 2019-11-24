package main

import (
	"./image_common"
	"log"
)

func main() {

	err := image_common.ResizeDir("/tmp/img/", 30, 60)
	if err != nil {
		log.Fatalln(err)
	}

}
