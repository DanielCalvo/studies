package main

import "fmt"

//Variadic functions can be called with any number of trailing arguments!
// https://en.wikipedia.org/wiki/Variadic_function
func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result

}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
