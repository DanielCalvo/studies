package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	a := 1

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				a++ //this looks ugly though, hmm
			}
		}
	}()

	for {
		fmt.Println(a) //value changes!
		time.Sleep(time.Second)
	}

	//ticker.Stop()
	//done <- true
	//fmt.Println("Ticker stopped")

}
