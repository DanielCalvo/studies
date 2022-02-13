package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	b, err := os.ReadFile("/home/daniel/Projects/studies/Golang/metryingtofigureshitout/files/counter_in_file.txt")
	if err != nil {
		fmt.Println(err)
	}
	bs := string(b)

	bi, err := strconv.Atoi(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bi)

	bi++
	bs = strconv.Itoa(bi)

	err = os.WriteFile("/home/daniel/Projects/studies/Golang/metryingtofigureshitout/files/counter_in_file.txt", []byte(bs), 0644)
	if err != nil {
		fmt.Println(err)
	}

}
