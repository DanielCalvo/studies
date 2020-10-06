package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Stat("/etc/passwdaaa")

	if err == nil {
		if file.IsDir() {
			fmt.Println("It's a directory")
		} else {
			fmt.Println("It's a file")
		}
	}
	if os.IsNotExist(err) {
		fmt.Println("That doesn't exist")
	} else {
		fmt.Println("Filesystem error:", err)
	}
}
