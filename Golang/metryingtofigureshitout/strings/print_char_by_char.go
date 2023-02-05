package main

import "fmt"

func main() {
	//Very basic, but I once forgot this.
	str := "Hello world"
	for _, v := range str {
		fmt.Print(string(v), " ")
	}
}
