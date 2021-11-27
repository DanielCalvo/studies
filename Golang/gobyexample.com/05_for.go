package main

import "fmt"

func main() {

	//basic
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//The classic initial/condition/after loop
	for a := 1; a <= 5; a++ {
		fmt.Println(a)
	}

	for {
		fmt.Println("Ever, until you break out of it!")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
