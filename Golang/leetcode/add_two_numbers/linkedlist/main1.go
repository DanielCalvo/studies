package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//Lets do something simple at first -- lets manually instantiate 3 elements and do some simple things with them
	l3 := ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := ListNode{
		Val:  2,
		Next: &l3,
	}
	l1 := ListNode{
		Val:  1,
		Next: &l2,
	}

	//fmt.Println(TranverseListAndSumValues(&l1))

	newHead := InsertItemInList(l1)
	PrintListItems(&newHead)

}
func PrintListItems(node *ListNode) {
	for node != nil {
		fmt.Println(*node)
		node = node.Next
	}
}

func InsertItemInList(head ListNode) ListNode {
	l := ListNode{
		Val:  -1,
		Next: &head,
	}
	return l
}

// Goes through the entire list if you pass the first node as an argument. Hey, lets sum the values!
func TranverseListAndSumValues(node *ListNode) int {
	var sum int
	//Remember -- you're checking if pointer to node is nil, not if node itself is nil. They're different things!
	for node != nil {
		sum += node.Val //if you put sum after node = node.Next, you can get an nil pointer, careful!
		node = node.Next
	}
	return sum
}

//How about adding an item at the end of the list?
