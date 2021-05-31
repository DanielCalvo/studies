package main

import (
	"fmt"
	"time"
)

/*
The goal of this is to understand how to have a dynamic number of workers: One worker per different domain!

But let's keep it simple for now:
Given a list of strings, have N number of workers based on the letter that they start with
*/

/*
First: Create a function to identify all the unique first letters that your strings begin with

for string in stringlist
	what letter does string begin with?
	add that string to a collection of things that begin by that character

*/

var thing = map[string][]string{}

func main() {

	mySlice := []string{"asdasd", "aaaaa", "aeoooo", "baswdasda", "bbbbbbbbbb", "bbfbdhfbdhfbdhfb"}

	for _, s := range mySlice {
		thing[string(s[0])] = append(thing[string(s[0])], s)
	}

	//you want to iterate over the keys on the map

	for k := range thing {
		ProcessElement(k) //Hmm, what arguments does this actually take?
	}
}

func ProcessElement(s *[]string) {
	fmt.Println("Processing element:", s)
	time.Sleep(time.Second)
}
