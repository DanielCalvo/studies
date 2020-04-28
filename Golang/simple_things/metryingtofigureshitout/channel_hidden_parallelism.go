package main

import (
	"fmt"
	"log"
	"sync"
)

func CheckStringsWorker(sIn <-chan string, workerNum int) <-chan string {
	sOut := make(chan string)
	var wg sync.WaitGroup

	wg.Add(workerNum)
	go func() {
		for i := 0; i < workerNum; i++ {
			go func() {
				defer wg.Done()
				for ss := range sIn {
					fmt.Println("On worker function:", ss)
					sOut <- ss
				}
			}()
		}
		wg.Wait()
		close(sOut)
	}()
	return sOut
}

func CheckStrings(s []string, workerNum int) []string {
	linkChan := make(chan string)
	go func() {
		for _, link := range s {
			linkChan <- link
		}
		close(linkChan)
	}()

	checkedStrings := CheckStringsWorker(linkChan, workerNum)

	var sProcessed []string

	for checkedLink := range checkedStrings {
		sProcessed = append(sProcessed, checkedLink)
		log.Println("Checked:", checkedLink)
	}
	return sProcessed
}

func main() {

	myStrings := []string{"banana", "apple", "orange"}
	ss := CheckStrings(myStrings, 8)
	for _, s := range ss {
		fmt.Println(s)
	}
}
