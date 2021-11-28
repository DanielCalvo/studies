package main

import (
	"errors"
	"fmt"
)

//Didn't finish this one!

func f1(n int) (int, error) { //By convention, errors are the last return type!
	if n == 42 {
		return -1, errors.New("Can't process 42!")
	}
	return n + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func main() {
	n, err := f1(2)
	if err != nil {
		panic("f1 crashed yo")
	}

	fmt.Println(n)
}
