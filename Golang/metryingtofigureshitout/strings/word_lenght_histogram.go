package main

import (
	"fmt"
	"strings"
)

func main() {

	//Hmm, I'm not sure if this is better than the C implementation. Interesting. I think a map would be better here perhaps?
	wordLenghts := make([]int, 50)
	myString := "Hello world, today is a good day! Longwooooooord"

	words := strings.Split(myString, " ")

	for _, word := range words {
		wordLenghts[len(word)]++
	}

	for k, v := range wordLenghts {
		if v != 0 {
			fmt.Println("There are", v, "words with a lenght of", k)
		}
	}

}
