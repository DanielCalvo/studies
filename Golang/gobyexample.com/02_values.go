package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)

	fmt.Println(true)
	fmt.Println(true && false) //Hey this is not true!
	fmt.Println(true || false) //True
	fmt.Println(!true)         //Not true, ha!
}
