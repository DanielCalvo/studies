package main

import (
	"fmt"
	"time"
)

func main() {

	d := "12/23/2019"                    //Why Americans, why...
	dt, err := time.Parse("1/2/2006", d) //Looks like this works? Kinda mysterious

	if err != nil {
		fmt.Println("Unable to time.Parse() provided date:", err)
	}

	fmt.Println(dt.Year(), dt.Month(), dt.Day())
	formattedDate := dt.Format("2006-01-02")
	fmt.Println(formattedDate)
}
