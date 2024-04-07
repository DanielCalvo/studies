package main

import (
	"errors"
	"fmt"
)

func main() {

	s := "42"
	n := 0
	for _, ch := range []byte(s) {
		ch = ch - '0' //Minus '0' actually means minus 48: https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html
		if ch > 9 {
			fmt.Println("Somethings wrong:", ch)
		}
		n = n*10 + int(ch) //And in here?
	}
	fmt.Println("n is:", n)

}

// Lets copy strconv.Atoi's and analyze it step by step

// This ends up being 64
const intSize = 32 << (^uint(0) >> 63)

func Atoi(s string) (int, error) {
	const fnAtoi = "Atoi"
	sLen := len(s)
	if intSize == 32 && (0 < sLen && sLen < 10) || intSize == 64 && (0 < sLen && sLen < 19) {
		// Fast path for small integers that fit int type.
		s0 := s
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0, errors.New(fnAtoi + s0) //Changed as importing the syntaxError function requires various dependencies
			}
		}

		n := 0
		//Hey this seems to be the interesting part
		for _, ch := range []byte(s) {
			ch -= '0'
			if ch > 9 {
				return 0, errors.New(fnAtoi + s0) //Changed as importing the syntaxError function requires various dependencies
			}
			n = n*10 + int(ch)
		}
		if s0[0] == '-' {
			n = -n
		}
		return n, nil
	}

	// Slow path for invalid, big, or underscored integers.

	// Hang on let me comment htis out for now
	//i64, err := ParseInt(s, 10, 0)
	//if nerr, ok := err.(*NumError); ok {
	//	nerr.Func = fnAtoi
	//}
	//return int(i64), err
	return 0, nil
}
