package main

import "fmt"

func main() {
	fmt.Println(wordCount("hello world"))
}

// lets figure out a way to count words as a start, lets start with a rudimentary approach first:
// lets do something incredibly stupid first, lets just count whitespace tranversing the string:
func wordCount(words string) (int, error) {
	return 0, nil
}
