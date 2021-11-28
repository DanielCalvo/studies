package main

import "fmt"

//range iterates over elements in a variety of data structures

func main() {
	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	//ranging over arrays and slices provides both an index and the value. You can ignore the index with the blank identifier if you don't need it
	for i, num := range nums {
		fmt.Println(i, num)

	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	//ranging over a map iterates over the key/value pairs
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	//ranging over a string iterates over unicode points
	//The first value is the starting byte index of the rune and the second is the rune itself
	for i, c := range "banana" {
		fmt.Println(i, c, string(c))
	}

}
