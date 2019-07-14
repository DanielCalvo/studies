package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {

	for {
		x, y := robotgo.GetMousePos()
		col := robotgo.GetMouseColor()
		fmt.Printf("\033[2K\r%d %d - %s", x, y, col)
		time.Sleep(200 * time.Millisecond)
	}

}
