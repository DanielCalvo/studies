package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sch := make(chan string)

	//This function will read the contents of /etc/passwd and send it to the sch channel in this go routine
	go sendStringToChannel(sch)

	//If we do this like this, the main go routine finishes before this gets a chance to run
	//go readStringFromChannel(sch)
	for s := range sch {
		fmt.Println(s)
	}

}

// Let's do a function that reads lines from a file and sends those to a channel of type string
// sendStringToChannel only sends data to the channel, so let's set the channel as send only on the function definition!
func sendStringToChannel(sch chan<- string) {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sch <- scanner.Text()
	}
	/*
		Very interesting, if I don't close the channel I get a deadlock in main
		Not all channels need to be closed, but you should close it when it is important
	*/
	close(sch)
	//return sch //
}

func readStringFromChannel(sch chan string) {
	for s := range sch {
		fmt.Println(s)
	}
}
