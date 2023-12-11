package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "000000010230"
	fmt.Println(RemoveZerosAtBeginning(s))
}

func ReverseNumber(num int) int {
	reversed := 0
	for num != 0 {
		remainder := num % 10
		reversed = reversed*10 + remainder
		num /= 10
	}
	return reversed
}

// hang on let me implement this in a really bad way I wanna have an adventure
// ABSOLUTE THRASH THIS IS MERELY ILLUSTRATIVE
func ReverseNumberString(num int) (int, error) {
	var is_negative bool
	s := strconv.Itoa(num)

	if num < 0 {
		is_negative = true
		s = s[1:]
	}
	s = ReverseString(s)
	//-1230 now you have: 0321

	//TERRIBLE, ABSOLUTELY TERRIBLE (uuuh remove all the 0s at the beginning of the string)
	s = RemoveZerosAtBeginning(s)

	if is_negative {
		s = "-" + s
	}

	//aahahah BUT WAIT you removed all zeroes above, can't convert an empty string to 0 wew lad
	if s == "" {
		s = "0"
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return num, nil
}

// DANG SO CLEAN SO SIMPLE THANKS STACK OVERFLOW :CHEFS-KISS:
// perhaps not the fastest, but it certainly is easy to read: https://stackoverflow.com/a/4965535
func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// Hey this function isn't that bad... if I was using this in the correct context
func RemoveZerosAtBeginning(s string) (result string) {
	counter := 0
	for _, v := range s {
		if string(v) == "0" {
			counter++
		} else {
			break
		}
	}
	return s[counter:]
}
