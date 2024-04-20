package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var LinkedListHead = ListNode{
		Val:  1,
		Next: nil,
	}

	PointerAddElementAtBeginning(&LinkedListHead, -20)

	fmt.Println(LinkedListHead) //Should be different

}

func PointerAddElementAtBeginning(l *ListNode, val int) {
	//l is LinkedListHead -- so just change the value?
	//l needs to point to another new node -- and this new node needs to point to the older instance of l
	newFirstElement := ListNode{
		Val:  val,
		Next: l,
	}
	*l = newFirstElement
}
