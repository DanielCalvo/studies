package main

import (
	"fmt"
)

func main() {

	//I found this in the strvconv package. What is going on here?
	intSize := 32 << (^uint(0) >> 63)
	fmt.Println(intSize) //Prints 64, uh-oh

	fmt.Println("---")
	fmt.Println(^uint(0)) //Creates an unsigned integer with all 64 bits set to 1, in decimal this is 18446744073709551615 (yep, 64 bits)
	//Lets doublecheck:
	fmt.Printf("Binary representation of ^uint(0): %b\n", ^uint(0)) //Yup, prints 1111111111111111111111111111111111111111111111111111111111111111

	//But what about this?
	//This performs a right shift by 63 bits on the value obtained in step 1
	fmt.Println(^uint(0) >> 63)                                               //prints 1 in decimal
	fmt.Printf("Binary representation of ^uint(0) >> 63: %b\n", ^uint(0)>>63) //Prints 1 in binary too

	//So this whole thing: 32 << (^uint(0) >> 63)
	//Is equivalent to: 32 << 1
	//This performs a left shift on 32 by 1
	fmt.Println(32 << 1) //prints 64, interesting

	//To illustrate, lets do some more left shifts:
	fmt.Println("Some left shifts:", 1<<1, 1<<2, 1<<3, 1<<4)
	fmt.Printf("Same left shifts in binary:: %b %b %b %b\n", 1<<1, 1<<2, 1<<3, 1<<4) //Prints 1 in binary too
	fmt.Println("---")
	//How about some right shifts?
	fmt.Println("Some right shifts:", 128>>1, 128>>2, 128>>3, 128>>4)
	fmt.Printf("Same right shifts in binary:: %b %b %b %b %b\n", 128>>0, 128>>1, 128>>2, 128>>3, 128>>4) //Prints 1 in binary too

}
