package main

import "fmt"

func main() {
	a := 11

	for i := 1; i < 100; i++ {
		a = a * i

	}
	fmt.Println(a)
}
