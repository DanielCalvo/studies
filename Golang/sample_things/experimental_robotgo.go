package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {

	//Takes and saves a screenshot, neat!
	bitmap := robotgo.CaptureScreen()
	fmt.Println("...", bitmap)
	robotgo.SaveBitmap(bitmap, "/home/daniel/PycharmProjects/studies/Golang/sample_things/test.png")

	//TODO next: Find a subimage of my screenshot. (feed another image into the program, find it inside the screenshot, get it's coordinates)

}
