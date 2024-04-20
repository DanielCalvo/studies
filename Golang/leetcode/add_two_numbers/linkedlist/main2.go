package main

import (
	"fmt"
)

// Okay now lets just create a single element and:
// 1. Add an item at the beginning at the list (with a function)
// 2. Add an item at the end of the list (with a function)
// 3. Make a function to pretty print the list (with a function!)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var LinkedListHead = ListNode{
		Val:  5,
		Next: nil,
	}

	LinkedListHead = AddElementAtBeginning(LinkedListHead, 4)
	LinkedListHead = AddElementAtBeginning(LinkedListHead, 3)
	LinkedListHead = AddElementAtBeginning(LinkedListHead, 2)

	//But how do I do this -- without returning an element, with pointers only?
	PrettyPrintList(&LinkedListHead)
	fmt.Println("Head:", LinkedListHead)

	AddElementAtEnd(&LinkedListHead, 55)
	AddElementAtEnd(&LinkedListHead, 66)

	//Hold up let me try something different
	var head *ListNode
	PointerAddElementAtBeginning(&head, 3)
	PointerAddElementAtBeginning(&head, 2)
	PointerAddElementAtBeginning(&head, 1)
	PrettyPrintList(head)
}

// I got this ** (pointer to pointer) shenanigan from chatGPT, I'm having a difficult time wrapping my head around the whole pointer to pointer thing entirely, but it does work
func PointerAddElementAtBeginning(l **ListNode, val int) {
	newNode := &ListNode{Val: val}
	newNode.Next = *l
	*l = newNode
}

// The above is very similar to what I was doing previously, but the below doesn't work for some reason, yet the function above does.
func AddElementAtBeginning(l ListNode, val int) ListNode {
	newFirstElement := ListNode{
		Val:  val,
		Next: &l,
	}
	return newFirstElement
}

func FindLastElement(l *ListNode) *ListNode {
	for l != nil {
		if l.Next == nil {
			return l
		}
		l = l.Next
	}
	return nil //if there's no last element (ex: its a circular list) then no luck
}

// You can't return the last element -- you gotta append it by pointer
func AddElementAtEnd(l *ListNode, val int) {
	Currentlast := FindLastElement(l) //You could have a circular list and this could return nil
	if Currentlast == nil {
		return
	}

	newLast := ListNode{
		Val:  val,
		Next: nil,
	}
	Currentlast.Next = &newLast
}

// Remember to receive a pointer here -- pointers can be nil, but the type itself can't
func PrettyPrintList(l *ListNode) {
	for l != nil { // There's a bug in here. The function will loop forever is the list is circular
		fmt.Print(*l)
		l = l.Next
	}
}
