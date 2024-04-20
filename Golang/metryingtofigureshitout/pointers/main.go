package main

import "fmt"

type Something struct {
	name   string
	number int
}

func main() {
	//One thing I realized working with linked lists: A struct can't be nil, but a pointer to a struct can be nil.

	//s := Something{}
	//s = nil you can't do this

	//but you can do this
	s1 := &Something{
		name:   "asd",
		number: 123,
	}
	fmt.Println(s1)
	fmt.Println(*s1)
	//You can assign nil to a pointer to something
	s1 = nil
	fmt.Println(s1) //prints nil

	if s1 != nil {
		fmt.Println(*s1)
	} else {
		fmt.Println("that's a nil pointer, trying to access a nil pointer during runtime causes go to panic")
	}

	//But hang on -- what about a pointer to a pointer?
	//Can you create a bunch of references to things that no longer exist? is this how you introduce a memory leak?
	var num int = 42
	ptr1 := &num
	fmt.Println(*ptr1)
	ptr2 := &ptr1
	fmt.Println(**ptr2)
	//So if I change ptr2's value, it should change num, and ptr1 too:
	**ptr2 = 55
	fmt.Println(num, *ptr1, **ptr2) //same value
	fmt.Println(&num, &ptr1, &ptr2) //different addresses
	//Can I be silly?
	ptr3 := &ptr2
	ptr4 := &ptr3
	ptr5 := &ptr4
	//Yeah I can be really silly
	fmt.Println(***ptr3, ****ptr4, *****ptr5)
	//Okay you can be stupidly silly: https://go.dev/play/p/GfplrZdO1cT

}
