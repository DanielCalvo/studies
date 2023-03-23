package main

import "fmt"

//Uh-oh, Go supports embedding structs and interfaces

type base struct {
	num int
}

// A container embeds a base
type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func main() {
	b := base{num: 4}
	fmt.Println(b.describe())

	c := container{
		base: base{
			num: 7,
		},
		str: "hello world",
	}
	fmt.Println(c)
	fmt.Println(c.base.num)

	//But wait, you can access the base fields directly!
	fmt.Println(c.num)
	//Since container embeds base, the methods of base also become methods of a container
	fmt.Println(c.describe())

	type describer interface {
		describe() string
	}

	//Embedding structs can be used to grant interface implementations onto other structs -- neat!
	var d describer = c
	fmt.Println(d.describe())

	//Note from self: I have a feeling this rabbit hole goes a bit deeper than what I touched on here!
}
