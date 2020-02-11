package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("sup")
	file, err := os.Open("./numbers.txt")
	check(err)

	counter := 0
	scanner := bufio.NewScanner(file)

	var intSlice []int
	innerCounter := 1

	for scanner.Scan() {
		counter++

		myInt, err := strconv.Atoi(scanner.Text())
		check(err)

		intSlice = append(intSlice, myInt)

		if counter == 100 {

			fmt.Println("counter reached", counter)
			sort.Ints(intSlice)
			fmt.Println(intSlice)
			//save to file
			f, err := os.Create(strconv.Itoa(innerCounter) + ".txt")
			if err != nil {
				fmt.Printf("error creating file: %v", err)
				return
			}
			defer f.Close()
			for _, num := range intSlice {
				_, err = f.WriteString(fmt.Sprintf("%d\n", num)) // not sure why there's a newline at the end of the file but ok
				if err != nil {
					fmt.Printf("error writing string: %v", err)
				}
			}
			counter = 0
			intSlice = nil
			innerCounter++
		}
	}
	fmt.Println("Last counterboi:", len(intSlice))
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	f, err := os.Create(strconv.Itoa(innerCounter) + ".txt")
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	for _, num := range intSlice {
		_, err = f.WriteString(fmt.Sprintf("%d\n", num)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
	//save to file

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
