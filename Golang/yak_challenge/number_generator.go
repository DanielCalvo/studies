package main

import (
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.OpenFile("./numbers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for i := 1; i <= 30; i++ {
		if _, err := f.WriteString(strconv.Itoa(i) + "\n"); err != nil {
			log.Println(err)
		}
	}

}
