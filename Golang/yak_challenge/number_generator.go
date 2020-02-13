package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	//10,000,000 (ten million) lines will generate a file that is 115mb long.
	//100,000,000 (a hundred million) lines will generate a file that is about 1gb long.
	//2,147,483,647 (two billion and something, biggest int32 possible) will generate a file that is about 21gb long.
	//Anything higher than that... Don't forget to bring your towel.

	var i int64

	filePath := flag.String("filepath", "/tmp/unsorted_list.txt", "Filesystem path to generate the unsorted list")
	numbers := flag.Int64("num", 10000, "Amount of numbers to generate on file")
	illegal := flag.Bool("illegal", true, "Would you like some non integer values on your file?")

	flag.Parse()

	//Quietly attempt to remove a previous list of random numbers on the default location
	_ = os.Remove(*filePath)

	f, err := os.OpenFile(*filePath, 0755, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i = 1; i <= *numbers; i++ {
		if _, err := f.WriteString(strconv.Itoa(rand.Intn(21477483647)) + "\n"); err != nil {
			log.Println(err)
		}
	}

	fmt.Println("Generated a file at", *filePath, "with", *numbers, "pseudorandom numbers")

	if *illegal {
		f.WriteString("I wonder if I'll get to find out what all the yak memes are about...\n")
		f.WriteString("123.123\n")
		f.WriteString("サウルは素晴らしい散髪をしています\n")
		f.WriteString("333 444\n")
		f.WriteString("42towel\n")
		f.WriteString("kidneys for sale, barely used, O+ blood type.\n") //Hey that's definitely illegal!
	}
}
