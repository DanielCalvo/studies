package main

import (
	"fmt"
)

func main() {
	//var x [5]int
	//x[3] = 42
	//fmt.Println(x)
	//fmt.Println(len(x))
	//Author says: We don't really use arrays in Go. Official documentation says: Use slices instead

	//A composite literal:
	// x := type{values}
	//a := []int{4,5,6,7,52}
	//fmt.Println(a)
	//fmt.Println(a[0])
	//fmt.Println(a[1:])
	//fmt.Println(a[:1])
	//
	//for i, v := range a {
	//	fmt.Println(i, v)
	//}
	//
	//for i := 0; i < len(a); i++ {
	//	fmt.Println("Position:", i, "Element:", a[i])
	//}
	//y := []int{123,33,213}
	//a = append(a,  y...)
	//
	//fmt.Println(a)
	//There is no delete function for slices in go. You use the append function to delete things!
	//Let's try removing the last element from our slice a.
	//a = append(a[:len(a)-1])
	//fmt.Println(a)

	//b := make([]int, 10, 100)
	//fmt.Println(b)
	//fmt.Println(len(b))
	//b = append(b, 111)
	//fmt.Println(b)
	//fmt.Println(len(b))
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(jb, mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)

	m := map[string]int{
		"James":           33,
		"Miss Moneypenny": 23,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Someone"]) //returns 0 if there is no entry

	if v, ok := m["banana"]; ok {
		fmt.Println("Printing inside the if")
		fmt.Println(v, ok)
	}
	if v, ok := m["James"]; ok {
		fmt.Println("Printing inside the if")
		fmt.Println(v, ok)
	}
	m["todd"] = 34
	fmt.Println(m["todd"])

	//Iterating through the map
	for k, v := range m {
		fmt.Println(k, v)
	}

	//Iterating through a slice
	xi := []int{4, 5, 6, 7, 8, 9, 42}

	for k, v := range xi {
		fmt.Println(k, v)
	}

	delete(m, "todd")
	for k, v := range m {
		fmt.Println(k, v)
	}

}
