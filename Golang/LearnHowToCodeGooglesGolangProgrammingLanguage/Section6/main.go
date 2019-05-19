package main

import (
	"fmt"
)

var x bool

func main() {
	a := 7
	b := 42

	fmt.Println(x)
	fmt.Println(a == b)
	fmt.Println(a != b)

	x := 42
	y := 42.312543
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	//fmt.Println(runtime.GOOS)
	//fmt.Println(runtime.GOARCH)

	s := "Hello world"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println()

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println()
	const e = 42
	const f = 42.78
	const g = "James bond"

	fmt.Println(e, f, g)
	fmt.Printf("%T %T %T", e, f, g)

	const (
		h = iota
		i = iota
		j = iota
	)
	fmt.Println(h, i, j)

	myvar1 := "banana"
	myvar2 := "apple"

	fmt.Println(myvar1 == myvar2)
	fmt.Println(myvar1 <= myvar2)
	fmt.Println(myvar1 >= myvar2)
	fmt.Println(myvar1 != myvar2)
	fmt.Println(myvar1 > myvar2)
	fmt.Println(myvar1 < myvar2)

}
