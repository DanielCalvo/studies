package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(DigitsInNumnber(0), DigitsInNumnber(77), DigitsInNumnber(-233), DigitsInNumnber(2313))
}

func DigitsInNumnber(i int) int {
	counter := 1 //Every number must have at least one digit in Go (I'm looking at you Javascript)
	for {
		if DivideBy10(i) == 0 {
			return counter
		}
		i = DivideBy10(i)
		counter++
	}
	return counter
}

func DivideBy10(i int) int {
	return i / 10
}

/*
Hmm some google searches and chatgpt returned this function using a conversion to string thing
I wanted to do this without using strings as I thought it would be faster. Lets benchmark this!
*/

func CountDigits(n int) int {
	absValue := int(math.Abs(float64(n)))
	numStr := strconv.Itoa(absValue)
	return len(numStr)
}
