package main

import "fmt"

func main() {

	a := "banana"
	//a[0] = '' //can't: empty rune literal or unescaped '
	//a[0] = 'e' //also can't do this -- strings are immutable in go
	fmt.Println(a[1])
	fmt.Println(a)

	//you can reassign the string in a variable, but not change the string, so
	aa := "apple"
	aa = "grape" //this is okay
	//aa[0] = e //this is not okay
	fmt.Println(aa)

	//Interesting!
	//A person on stackoverflow says: Strings are immutable in Go

	r := []rune(a)
	r[0] = rune('0')
	//r[0] = rune('') //can't throw and empty literal in there either
	fmt.Println(r, " - ", string(r))
}
