package main

import (
	"fmt"
	"os"
)

//Echo prints it's command line arguments

func main() {
	fmt.Println(os.Args[1:])

	fmt.Println(os.Args)

	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}
