package main

import (
	"fmt"
	"time"
)

type numbersito struct {
	Number    int
	SqrNumber int
}

func main() {

	//3 channels
	//new = instantiates a new element and pushes it into a channel
	//sqrt = receives element from channel, square roots it, pushes it into another channel. Fan out here!
	//finalize = prints squared element and exits. Fan in here!

	in := genNumbersito(5)
	c1 := sqrtNumbersito(in)
	for n := range c1 {
		fmt.Println(n)
	}

}

func genNumbersito(baseNum int) <-chan numbersito {
	out := make(chan numbersito)
	go func() {
		for i := baseNum; i < baseNum+100; i++ {
			nu := numbersito{
				Number: i,
			}
			out <- nu
		}
		close(out)
	}()
	return out
}

func sqrtNumbersito(in <-chan numbersito) <-chan numbersito {
	out := make(chan numbersito)
	go func() {
		for n := range in {
			n.SqrNumber = n.Number * n.Number
			out <- n
		}
		close(out)
	}()
	return out
}

func printNumbersito(in <-chan numbersito) {
	go func() {
		for n := range in {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(n)
		}
	}()
}
