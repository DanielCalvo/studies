package main

import "fmt"

func main() {

	md := `### Hello!
You have opened dummy.md. I hope you have a good day!
- one item
 - one item that is not well indented :o`

	fmt.Println("before")
	fmt.Println(md)
	fmt.Println("after")

}
