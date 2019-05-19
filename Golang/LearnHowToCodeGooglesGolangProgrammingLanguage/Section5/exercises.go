package Section5

import "fmt"

//Exercise 2
var a int
var b string
var c bool

//Exercise 3
var e int = 42
var f string = "James Bond"
var g bool = true

//Exercise 4
type dani int

var xx dani

//Exercise 5
var yy int

//The short declaration operator can only be used at the function level. Can't be used at the package level

func main() {

	//Exercise 1
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)

	//Exercise 2
	//Prints 0, empty string, false. Those are the "zero values" for those types
	fmt.Println(a, b, c)

	//Exercise 3
	fmt.Sprintf("%v %v %v", e, f, g)
	s := f
	fmt.Println(s)

	//Exercise 4
	fmt.Printf("%T\n", xx)
	fmt.Println("Value of x: ", xx)
	xx = 42
	fmt.Println("Value of x: ", xx)

	//Exercise 5
	yy = int(xx)
	fmt.Println("Converted value: ", yy)

}
