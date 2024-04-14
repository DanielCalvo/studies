package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	s := "425"
	n := 0
	for _, ch := range []byte(s) {
		ch = ch - '0' //Minus '0' actually means minus 48: https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html
		if ch > 9 {
			fmt.Println("Somethings wrong:", ch)
		}
		n = n*10 + int(ch) //And in here?
	}
	fmt.Println("n is:", n)

	//Hang on, let me do single string character to int for now

	fmt.Println(Atoi("- 22"))

}

// One improvement: Could've used unicode.IsDigit
func Atoi(s string) int {
	isNegative := false
	result := 0
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	//if the first element is not a sign or a number, return
	if s[0] != '-' && s[0] != '+' && !IsDigit(s[0]) {
		return 0
	}

	if s[0] == '-' {
		isNegative = true
	}

	//once you find a sign or number, the next thing can only be a number, otherwise return result

	//Maybe you don't need this anymore and can rever to:
	//for _, c := range []byte(s) {
	for i := 0; i < len(s); i++ {
		c := s[i] //Careful using s[i] as s[i] is different from c after you do c - '0'

		if IsDigit(c) {
			c = c - '0'
			result = result * 10
			result = result + int(c)
			if result > math.MaxInt32 { //result is comparing to the wrong thing in here, 52 instead of 4
				if isNegative {
					return -math.MaxInt32 - 1
				} else {
					return math.MaxInt32
				}
			}
		}
		//if the next element is not a number, stop processing (but first check if its not out of bounds)
		if i+1 < len(s) {
			if !IsDigit(s[i+1]) {
				break
			}
		}
	}

	if isNegative {
		result = -result
	}

	return result
}

func IsDigit(b byte) bool {
	b = b - '0' //Minus '0' actually means minus 48: https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html
	if b <= 9 {
		return true
	}
	return false
}

// Lets copy strconv.Atoi's and analyze it step by step

// This ends up being 64
const intSize = 32 << (^uint(0) >> 63)
