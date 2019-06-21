package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {

	//xfce_menu_bitmap := robotgo.OpenBitmap("/home/daniel/PycharmProjects/studies/Golang/sample_things/xfce_start.png")
	//Move mouse to a part of the screen:
	//xfce_x, xfce_y := robotgo.FindBitmap(xfce_menu_bitmap, robotgo.CaptureScreen(), 0.1)
	//robotgo.MoveMouseSmooth(xfce_x, xfce_y)

	for {
		x, y := robotgo.GetMousePos()
		fmt.Println(x, y)
		time.Sleep(10)
	}

	//TODO next: Find a subimage of my screenshot. (feed another image into the program, find it inside the screenshot, get it's coordinates)

}
