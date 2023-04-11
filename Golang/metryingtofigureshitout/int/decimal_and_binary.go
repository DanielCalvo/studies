package main

import "fmt"

func notmain() {

	var a int = 0b1110011
	fmt.Println(a)        //prints it in decimal
	fmt.Printf("%b\n", a) //prints it in binary
	var aa int = 115

	//but hang on they're both integers so... they're the same thing right?
	if a == aa {
		fmt.Println("yeah its the same deal")
	}

	//interesting!
	fmt.Printf("%b --- %b\n", 100, -100)

}
