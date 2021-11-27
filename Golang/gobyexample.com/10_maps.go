package main

import "fmt"

func main() {

	//Maps are Go's associative data type: https://en.wikipedia.org/wiki/Associative_array
	m := make(map[string]int)
	fmt.Println(m)

	m["banana"] = 1
	m["apple"] = 2
	m["grape"] = 3
	fmt.Println(m["banana"])
	fmt.Println(len(m))

	delete(m, "apple")
	fmt.Println(m["apple"]) //zeroed

	nm := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(nm)

}
