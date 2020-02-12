package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	//hey I should implement some parallelism in here!
	//receive some arguments!

	//10,000,000 (ten million) lines will generate a file that is 115mb long.
	//100,000,000 (a hundred million) lines will generate a file that is about 1gb long.
	//2,147,483,647 (two billion and something, biggest int32 possible) will generate a file that is about 21gb long.

	//Anything higher than that... Don't panic and don't forget your towel.

	//Don't forget to add invalid input too!
	//print help message if ran with no arguments!

	var lines int64
	var i int64
	lines = 42424

	f, err := os.OpenFile("/tmp/numbers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i = 1; i <= lines; i++ { //A million!
		if _, err := f.WriteString(strconv.Itoa(rand.Intn(21477483647)) + "\n"); err != nil {
			log.Println(err)
		}
	}
	fmt.Println("Done")

}
