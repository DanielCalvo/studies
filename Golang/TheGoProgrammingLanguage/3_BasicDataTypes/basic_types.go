package main

import "fmt"

func main() {

	fmt.Println("Bitwise AND:", 2&2)
	fmt.Println("Bitwise OR:", 32|32, 2|4) //Bitwise OR
	fmt.Println("Bitwise XOR:", 128^128, 32^64)
	fmt.Println("Bitwise AND NOT:", 64&^64, 64&^32)
	fmt.Println("Left shift:", 2<<4)
	fmt.Println("Right shift:", 2>>4)

	fmt.Printf("%08b", 1<<1|1<<5) //Shift a bit in the first and fifth positions of.. a byte?

	fmt.Println()
	s := "Hello world"
	fmt.Println(s[1])

}
