package main

import "fmt"

// lets create something that satisfies the Writer interface with some dummy implementation!
type MyWriter struct{}

func (MyWriter) Write(p []byte) (n int, err error) {
	fmt.Println("could have done something useful here! p is:", p)
	return len(p), nil
}

func main() {
	s := "aaaaa"
	mw := MyWriter{}
	fmt.Fprintln(mw, s) //noice, looks like this handles the string to slice of byte type conversion!
}
