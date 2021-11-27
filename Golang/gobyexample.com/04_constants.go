package main

import (
	"fmt"
	"math"
)

const s string = "This is a constand of string type"

var a string = "And this is a global variable!"

func main() {
	fmt.Println(s)
	fmt.Println(a)

	const d = 3 //A numeric constant has no type until it is given one, interesting!

	fmt.Println(int64(d))
	fmt.Println(math.Sin(d)) //A constant number can be given a number by using it in a context that requires one

}
