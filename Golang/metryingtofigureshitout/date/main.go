package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	d := "1/23/2014" //Why Americans, why...
	//layout := "2006-30-03"
	fmt.Println("heyaaa")
	dt, err := time.Parse(d, d)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dt)
}
