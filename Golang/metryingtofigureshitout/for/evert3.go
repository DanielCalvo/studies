package main

import (
	"bufio"
	"fmt"
	"os"
)

//This is probably silly and/or very basic, but I want to have a reference of this to copy and paste around

func main() {

	f, _ := os.Open("/etc/passwd")

	scanner := bufio.NewScanner(f)

	counter := 0

	for scanner.Scan() {
		counter++
		if counter%3 == 0 {
			fmt.Println(counter)
		}
	}
	fmt.Println("Total lines:", counter)

}
