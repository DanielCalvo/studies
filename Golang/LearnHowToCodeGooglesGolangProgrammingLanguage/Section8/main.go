package main

import (
	"fmt"
)

func main() {

	//for init; condition; post
	//no while loop in go

	//for i := 0; i <= 100; i++ {
	//	fmt.Println(i)
	//}
	//for i := 0; i <= 5; i++ {
	//	for j := 0; j <= 5; j++ {
	//		fmt.Println("Outer loop:", i, "Inner loop:", j)
	//	}
	//}

	//x := 1
	//for x < 10 {
	//	fmt.Println(x)
	//	x++
	//}

	//a := 0
	//for {
	//	fmt.Println(a)
	//	a++
	//	if a > 5 {
	//		fmt.Println("breaking")
	//		break
	//	}
	//}

	//x := 1
	//for {
	//	x++
	//	if x > 100 {
	//		break
	//	}
	//	if x%2 != 0 {
	//		continue
	//	}
	//	fmt.Println(x)
	//}

	//loop over from the numbers 0 to 200 and print them out as ASCII text
	//s := "Hello world"
	//bs := []byte(s)
	//
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%#U ", s[i])
	//}
	//fmt.Println()

	//for i := 65; i <= 122; i++ {
	//	//s := string(i)
	//	//fmt.Println(i, s)
	//	fmt.Printf("%v %#U\n", i, i)
	//}
	if true {
		fmt.Println("Always true")
	}
	if x := 42; x == 2 {
		fmt.Println("True!")
	} else if x == 1 {
		fmt.Println("x equals to one")
	} else {
		fmt.Println("Nothing matched!")
	}

}
