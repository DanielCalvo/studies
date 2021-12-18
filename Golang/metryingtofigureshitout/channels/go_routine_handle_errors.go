package main

import (
	"errors"
	"fmt"
	"time"
)

//Incomplete!

func main() {
	go func() {
		//do something that returns an error
		//handle that error in  main!
		//I think I just need to push into a channel in here and pull from the channel in main and that's about it!
		err := errors.New("Hey something's wrong!")
		if err != nil {
			panic(err) //This is bad!
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Hey I reached the end of main!")
}
