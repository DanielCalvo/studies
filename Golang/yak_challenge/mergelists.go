package main

import (
	"fmt"
	"os"
)

func checkk(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("sup")
	file, err := os.Open("numbers.txt")
	checkk(err)
	o2, err := file.Seek(6, 0)
	checkk(err)

	b2 := make([]byte, 2)
	n2, err := file.Read(b2)
	checkk(err)

	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

}
