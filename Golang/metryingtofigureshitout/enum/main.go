/*
If I want fields of a struct to be only of a certain type, like imagine I want a string to represent the seasons in a year
Imagine it can only be winter, summer, spring or fall, how do I do that?
I remember Rust has enums, but how can I do this in Go?

Hmm, google returned this: https://www.educative.io/answers/what-is-an-enum-in-golang
*/

package main

import "fmt"

type Season int

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

func (s Season) toString() string {
	switch s {
	case Spring:
		return "Spring"
	case Winter:
		return "Winter"
	case Summer:
		return "Summer"
	case Autumn:
		return "Fall"
	default:
		return "invalid"
	}
}

func main() {
	a := Season(Winter)
	fmt.Println(a.toString())

}
