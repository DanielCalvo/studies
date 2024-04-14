package main

import "fmt"

func main() {
	//Very basic, but I once forgot this.
	str := "Hello world"
	for _, v := range str {
		fmt.Print(string(v), " ")
	}

	fmt.Println()
	//With the classic C loop
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
	}

	//lets also do the classic loop

}
