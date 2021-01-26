package main

import "fmt"

//https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go
func ThreeFirst(si []int) []int {
	var tmpSlice []int

	for i, v := range si {

		if v == 3 {
			fmt.Println("Element we want at position:", i)
			tmpSlice = append(tmpSlice, v)
			//si = append(si[:i-1], si[i:]...)
			fmt.Println(si[:i], si[i:])

		}
	}
	si = append(tmpSlice, si...)
	return si
}

func main() {

	xi := []int{1, 2, 3, 4, 5}
	//xi = xi[2:]

	//element to be removed is on index 2
	fmt.Println(xi[2:], xi[3:])
	copy(xi[2:], xi[3:])
	fmt.Println(xi[2:], xi[3:])

	//xi[len(xi)-1] = 0     // Erase last element (write zero value).
	//fmt.Println(xi)
	//xi = xi[:len(xi)-1]     // Truncate slice.
	//fmt.Println(xi)

	//var xa []int
	//for _, v := range xi {
	//	if v == 3 {
	//		xa = append(xa, v)
	//	}
	//}
	//for _, v := range xi {
	//	if v != 3 {
	//		xa = append(xa, v)
	//	}
	//}
	//fmt.Println(xa)
}
