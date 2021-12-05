package main

import "fmt"

func multiply(a, b int) int {
	fmt.Println("You called the multiply function!")
	return a * b

}

func main() {
	myvariable := multiply(2, 5)
	fmt.Println(myvariable)
	fmt.Println("The program finished!")
}
