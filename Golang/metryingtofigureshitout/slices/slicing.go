package main

import "fmt"

func main() {
	slicerino := []int{1, 2, 3, 4, 5}
	fmt.Println(slicerino)

	fmt.Println("Slicerino:", slicerino)
	fmt.Println("Exclude first element:", slicerino[1:])
	fmt.Println("Exclude last element:", slicerino[:len(slicerino)-1])
	fmt.Println("Exclude first and last element:", slicerino[1:len(slicerino)-1])
	fmt.Println("Only first element:", slicerino[0])
	fmt.Println("Only last element:", slicerino[len(slicerino)-1])
}
