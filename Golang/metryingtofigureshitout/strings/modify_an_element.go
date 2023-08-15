package main

import "fmt"

func main() {

	a := "banana"
	//a[0] = '' //can't: empty rune literal or unescaped '
	//a[0] = 'e' //also can't do this!
	fmt.Println(a[1])
	fmt.Println(a)

	//Interesting!
	//A person on stackoverflow says: Strings are immutable in Go

	r := []rune(a)
	r[0] = rune('0')
	//r[0] = rune('') //can't throw and empty literal in there either
	fmt.Println(r, " - ", string(r))
}
