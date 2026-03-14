package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(BareReturn())

}

// this is a bare return, and you can have it when you have named return values
func BareReturn() (n int, e error) { //go only allows this if you have named return values
	n = 11
	e = errors.New("oh no")
	return //retuns n and e -- best used sparingly as it can make the code hard to read
}

//this does not work, return values need to be named
//func WrongBareReturn() (int, error) {
//	n := 11
//	e := errors.New("oh no")
//	return
//}
