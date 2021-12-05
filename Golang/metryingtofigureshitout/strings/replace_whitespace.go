package main

import (
	"fmt"
	"strings"
)

func main() {

	myString := "   - This is a item on a markdown list that starts with too much whitespace! Let's trim it!"
	fmt.Println(strings.Trim(myString, " "))

	otherString := "- I want to add more whitespace to the beginning of this one!"
	fmt.Println(otherString)

}
