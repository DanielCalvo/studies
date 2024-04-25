package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := 1234
	List := CreateListWithElements(s)
	PrettyPrintList(List)

}

// ChatGPT helped me with this one. Wow. So simple and so smart. I could've never figured this out on my own with my beginner skills. How cool!
func CreateListWithElements(num int) *ListNode {
	var head *ListNode
	var prev *ListNode
	num = ReverseNumber(num)
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

func PrettyPrintList(l *ListNode) {
	for l != nil { // There's a bug in here. The function will loop forever is the list is circular
		fmt.Println(*l)
		l = l.Next
	}
}

func ReverseNumber(i int) int {
	var result int
	for i != 0 { //FIRST TRY, YEEEEEEEEAH (ok it was second try, I messed up the for loop, but still YEAAAH)
		result *= 10
		result += i % 10
		i = i / 10
	}
	return result
}
