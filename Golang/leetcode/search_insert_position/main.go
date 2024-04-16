package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{2, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))

}

// What if I get a negative target?
// O(log n) runtime complexity: <- You did not research/meet this target!
// You could make this a lot faster by only looking at parts of the array instead of going through the whole thing!
func searchInsert(nums []int, target int) int {
	var k, v int
	for k, v = range nums {
		if target == v {
			return k
		}
		if v > target {
			return k
		}
	}
	return k + 1
}
