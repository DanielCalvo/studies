package main

import "fmt"

func main() {

	b := make([]byte, 5)

	//I didn't think assigning a character to the element of a slice of byte would be valid!
	b[0] = 'a'
	b[1] = 97 //a

	if b[0] == b[1] {
		fmt.Println("amahgaaad they're the same")
	}

}
