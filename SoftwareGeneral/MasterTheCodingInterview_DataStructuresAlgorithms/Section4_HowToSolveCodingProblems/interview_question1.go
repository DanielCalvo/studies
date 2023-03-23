package main

import "fmt"

/*
Given 2 slices, create a function that lets a user know (true/false) whether these two arrays contain any common items
For example
slice1 := ['a', 'b', 'c']
slice2 := ['z', 'y', 'i']
should return false!

slice1 := ['a', 'b', 'c', 'x']
slice2 := ['z', 'y', 'x']
Should return true!
*/

/*
Your observations: Are the slices organized alphabetically?
*/

func main() {
	slice1 := []string{"a", "b", "c", "x"}
	slice2 := []string{"z", "y", "k"}

	//fmt.Println(CommonItemsSlow(slice1, slice2))
	fmt.Println(CommonItemsBetter(slice1, slice2))

}

// O(a*b) -- Quickest solution to write, but does not perform well. Brute force approach!
func CommonItemsSlow(s1, s2 []string) bool {
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

// This is a better solution for time complexity
func CommonItemsBetter(s1, s2 []string) bool {
	m := make(map[string]bool)

	for _, v1 := range s1 {
		if _, ok := m[v1]; !ok { //You only need to set this to true once per element
			m[v1] = true
		}
	}

	for _, v2 := range s2 {
		if m[v2] == true {
			fmt.Println("Match!")
			return true
		}
	}

	fmt.Println(m)
	return false
}
