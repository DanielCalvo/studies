package main

import "fmt"

// Iota is a constant generator, which is used to create a sequence of related values without spelling out each one individually
type Weekday int

const (
	Sunday  Weekday = iota //0
	Monday                 //1
	Tuesday                //2, and so on
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota //This evaluates to successive powers of two. But these are given a boolean interpretation, uh-oh
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	var v Flags = FlagUp
	fmt.Printf("%b %t\n", v, isUp(v)) // "10001 true"

}

func isUp(v Flags) bool {
	return v&FlagUp == FlagUp
}
