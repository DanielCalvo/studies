package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Thanks wikipedia, very cool:
//https://en.wikipedia.org/wiki/External_sorting
//https://en.wikipedia.org/wiki/Merge_algorithm#K-way_merging

//make the upper limit a variable

//I think I should put Yak references on the code in hopes they think I'm cool...

func main() {
	fmt.Println("It's yak time!1!one")

	//file := os.Args[1]
	//number, err := strconv.Atoi(os.Args[2])
	//if err != nil {
	//	log.Panic("Could not convert" + os.Args[2] + "to int")
	//}
	//
	//fmt.Println(file, number)
	//
	//_, err = os.Stat(file)
	//
	//if os.IsNotExist(err) {
	//	log.Panic("ERROR: input file does not exist.")
	//}
	//
	//if os.IsPermission(err) {
	//	log.Panic("ERROR: input file is not readable.")
	//}
	//
	//if number <= 0 {
	//	log.Panic("ERROR: Number of top results must be bigger than 0.")
	//}
	//
	//if number <= 30000000 {
	//	log.Panic("ERROR: Maximum number of top results must be less or equal than 30000000.")
	//}

	file, err := os.Open("./numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//read 10 lines
	//sort them
	//write them to a file
	//write many files (lists!)

	//k-way merge
	//get the first element of all the already sorted lists
	//of all these elements, write the smallest one to the final list
	//remove this element from the list it came from, or move the cursor forward
	//repeat the process until you run out of elements on all the lists and you're done!

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
