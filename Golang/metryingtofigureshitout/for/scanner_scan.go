package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal("Canot open /etc/passwd")
	}
	scanner := bufio.NewScanner(f)
	counter := 0

	//Kinda basic, but I always forget this
	for scanner.Scan() {
		counter++
	}
	fmt.Println("Total lines:", counter)

}
