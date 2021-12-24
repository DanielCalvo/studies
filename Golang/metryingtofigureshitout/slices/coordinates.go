package main

import "fmt"

func main() {

	coordinates := make([][]int, 2)

	for i := 0; i < len(coordinates); i++ {
		coordinates[i] = make([]int, 2)
	}
	coordinates[0][0] = 1

	fmt.Println(coordinates)

}
