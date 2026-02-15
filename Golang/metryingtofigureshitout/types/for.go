package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	//when iterating over a slice, v is always a copy of the element!
	for _, v := range nums {
		v = v * 10 // Only the copy is changed -- the elements in the slice remain the same!
	}
	fmt.Println(nums) // [1 2 3], the slice is unchanged!

	//if you want to change the slice, you need to update it by changing the indexed value
	for i := range nums {
		nums[i] = nums[i] * 10 //writes to the actual element
	}
	fmt.Println(nums) //[10 20 30], the slice is updated

}
