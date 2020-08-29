package main

import (
	"fmt"
	"time"
)

//program is more like
//run loop, time loop execution
//if more than 50 seconds have passed, call function to get new token


func main(){
	//var mynum int
	//mynum = 10
	//fmt.Println(mynum)

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	//var i int
	//for i=0; i<5; i++{
	//	time.Sleep(1)
	//	fmt.Println(i)
	//}

}
