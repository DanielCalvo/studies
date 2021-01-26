package main

import (
	"fmt"
	"sync"
)

func main() {

	//How to make main wait for a go routine to finish?
	var wg sync.WaitGroup
	wg.Add(1)
	go dostuff(&wg)
	wg.Wait()

}

func dostuff(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
