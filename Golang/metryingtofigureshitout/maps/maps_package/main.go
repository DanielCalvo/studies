package main

import (
	"fmt"
	"golang.org/x/exp/maps"
)

func main() {

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)

	fmt.Println(maps.Keys(m))
	fmt.Println(maps.Values(m))
	//Package has a few other useful functions. Nothing groundbreaking, but pretty cool nonetheless!
}
