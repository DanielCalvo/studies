package main

import "fmt"

//This program uses empty channels to block the go routines until the main function pulls empty values from it. Very cool!

func main() {
	ch := make(chan struct{})
	numbers := []int{4, 5, 6}

	for _, n := range numbers {
		go func(n int) {
			for i := 0; i < n; i++ {
				fmt.Println(i)
			}
			ch <- struct{}{}
		}(n)
	}

	//We launched one go routine per element of numbers, so we loop thorugh every element of number and remove something from the channel
	for range numbers {
		<-ch
	}

}

func countToSomething(a int) {
	for i := 0; i < a; i++ {
		fmt.Println(i)
	}
}
