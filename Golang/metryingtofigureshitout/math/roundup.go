package main

import (
	"fmt"
)

func main() {

	//This is really stupid but I don't remember how to do this
	//Round everything up to the closest multiple of four
	fmt.Println(roundup(11, 4))

}

func roundup(numToRound int, multiple int) int {
	remainder := numToRound % multiple
	return numToRound + multiple - remainder
}
