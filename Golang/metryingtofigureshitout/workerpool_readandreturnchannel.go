//write a function that will launch 5 workers to pull elements from a channel that it will receive as an argument
//do something the value from that channel (like multiply by 2)
//the function returns a channel which will receive the multiplied by 2 values

package main

import (
	"log"
	"sync"
	"time"
)

func genNumbersito(baseNum int) <-chan int {
	out := make(chan int)
	go func() {
		for i := baseNum; i < baseNum+10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func bytwo(num <-chan int, workerNum int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(workerNum)
	go func() {
		for i := 0; i < workerNum; i++ {
			go func() {
				defer wg.Done()
				for n := range num {
					time.Sleep(time.Millisecond * 200)
					out <- n * 2
				}
			}()
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {

	numbersitos := genNumbersito(40)
	twonumbersitos := bytwo(numbersitos, 3)

	for n := range twonumbersitos {
		log.Println(n)
	}
}
