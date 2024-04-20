//Chatgpt sent me this -- I want to finish understanding what's going on in here, it seems that pointer to pointer, while apparently not very common, can be useful

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var head *ListNode
	PointerAddElementAtBeginning(&head, 3)
	PointerAddElementAtBeginning(&head, 2)
}

func PointerAddElementAtBeginning(l **ListNode, val int) {
	newNode := &ListNode{Val: val}
	newNode.Next = *l
	*l = newNode
}
