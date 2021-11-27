package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd") //Who would've thought...
	}

	if num := 19; num < 0 { //A statement can precede conditionals!
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits") //Hey but what if it's negative and has multiple digits?
	}

}
