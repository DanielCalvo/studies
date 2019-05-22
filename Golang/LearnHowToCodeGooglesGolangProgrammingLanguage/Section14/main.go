package main

import "fmt"

func main() {
	foo()
	bar("Banana")
	s1 := woo("Apple")
	fmt.Println(s1)
	x, y := mouse("Dani", "Verydani")
	fmt.Println(x, y)
	a := sum(1, 2, 3, 4, 5)
	fmt.Println(a)
	xi := []int{1, 2, 3, 4, 5}
	sum(xi...)
	fmt.Println()
	defer foo()
	bar("asd")

}

//func (r receiver) identifier (parameters) (return) { ... }

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello from bar", s)
}

func woo(s string) string {
	return s
}

func mouse(fn string, ln string) (string, bool) {
	fmt.Println("Hello from mouse!", fn, ln)
	return fn, true
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
	return sum
}
