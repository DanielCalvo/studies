package main

import "fmt"

func main() {

	//can you have a negative byte?
	//no, you cant, byte is an uint8 (aka unsigned integer with 8 bits, ranges from 0 to 255)
	var b byte

	b = 2
	fmt.Println(b)
	//b = -1 //can't do this

	//But I can overflow it!
	b = b - 10
	fmt.Println(b) //prints 248, oof
}
