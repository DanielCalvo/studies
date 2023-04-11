package main

import "fmt"

func main() {
	fmt.Println("Heya")

	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(11))

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	num := x
	for num != 0 {
		remainder := num % 10
		reversed = reversed*10 + remainder
		num /= 10
	}
	//You don't need an if statement, this is a boolean thing
	return x == reversed

}
