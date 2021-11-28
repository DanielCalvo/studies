package main

import "fmt"

//Go supports anonymous functions!
// https://en.wikipedia.org/wiki/Anonymous_function

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	myfunc := intSeq()

	//myfunc retains state -- interesting!
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())

}
