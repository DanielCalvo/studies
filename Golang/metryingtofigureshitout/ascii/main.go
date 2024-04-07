package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Lets print all those out: https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html

	//Hey look, new for loop on go 1.22:
	//for i := range 127 { //ehhh but this is not what I want in the end

	for i := 32; i <= 127; i++ {
		fmt.Println("i: ", i, string(i))
	}

	//Hey hang on can I do it the other way round?
	s := "0123456789"
	fmt.Println("The type of a string element is: ", reflect.TypeOf(s[0])) //This is of type uint8
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) //Easier than I thought!
	}

}
