package main

import "fmt"

func main() {

	x := 42
	func(x int) {
		fmt.Println("The meaning of life is:", x)
	}(x)

}
