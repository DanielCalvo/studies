package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("I wonder what happens!")
	go bananas()

	time.Sleep(time.Second)
	fmt.Println("End of main")
}

func bananas() {
	fmt.Println("I have gone bananas!")
	panic("bananas") //This panics the entire program, not just the go routine
}
