package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//Okay but how I do go about solving the actual problem now...
	//The digits in the list are stored in reverse order

	n3 := ListNode{
		Val:  3,
		Next: nil,
	}
	n2 := ListNode{
		Val:  4,
		Next: &n3,
	}
	n1 := ListNode{
		Val:  2,
		Next: &n2,
	}

	nn3 := ListNode{
		Val:  4,
		Next: nil,
	}
	nn2 := ListNode{
		Val:  6,
		Next: &nn3,
	}
	nn1 := ListNode{
		Val:  5,
		Next: &nn2,
	}

	PrettyPrintList(addTwoNumbers(&n1, &nn1))
	fmt.Println()

	var l1 ListNode
	var l2 ListNode
	l3 := addTwoNumbers(&l1, &l2)
	PrettyPrintList(l3)

}

func NumberToList(num int) *ListNode {
	var head *ListNode
	var prev *ListNode
	//num = ReverseNumber(num)
	for num != 0 { //Its doing it in reverse but lets roll with it for now

		//fmt.Println("This would be added to a list element:", num%10)
		node := &ListNode{Val: num % 10}

		if head == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
		num = num / 10
	}
	return head
}

// If list contains 123, desired number is 321
func GetNumberFromList(l *ListNode) int {
	//Get number and then apply the "reverse number" trick?
	var num int

	//from the "printpretty" function
	for l != nil { // There's a bug in here. The function will loop forever is the list is circular
		num *= 10
		num += l.Val
		l = l.Next
	}
	//W-w-wait but the problems are
	num = ReverseNumber(num)
	return num
}

// Receives 123, returns 321
func ReverseNumber(i int) int {
	var result int
	for i != 0 { //FIRST TRY, YEEEEEEEEAH (ok it was second try, I messed up the for loop, but still YEAAAH)
		result *= 10
		result += i % 10
		i = i / 10
	}
	return result
}

// Get both numbers, sum them, add them to a list?
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := GetNumberFromList(l1)
	num2 := GetNumberFromList(l2)
	fmt.Println(num1, num2)

	result := num1 + num2
	fmt.Println("Result:", result)

	var list *ListNode
	list = NumberToList(result)
	return list
}

// Doesn't work if you pass a list with a single element
func PrettyPrintList(l *ListNode) {
	for l != nil { // There's a bug in here. The function will loop forever is the list is circular
		fmt.Print(l.Val, " -> ")
		l = l.Next
	}
}
