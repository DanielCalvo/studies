package main

import (
	"fmt"
	"time"
)

func main() {

	// date +"%Y-%m-%dT%H:%M:%S%:z"

	//Time from go: 2020-04-20T11:44:11+02:00
	dateFromShell := "2020-04-20T11:45:03+02:00"

	timeDate, err := time.Parse(time.RFC3339, dateFromShell)
	fmt.Println(err)
	fmt.Println(timeDate)

	fmt.Println("Time now:")
	fmt.Println(time.Now().Format(time.RFC3339))

}
