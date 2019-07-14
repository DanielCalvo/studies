package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	//Manabar pos 1024 32

	fmt.Println(robotgo.GetPixelColor(1024, 32))
	//when full: 00469b
	//when empty: 2a2b2a
	//something else: Tibia not open or not allignet
}
