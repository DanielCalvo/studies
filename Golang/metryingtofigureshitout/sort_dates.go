package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	now := time.Now().Format(time.RFC3339)
	fmt.Println(now)

	nowString := string(now)
	fmt.Println(nowString)

	nowAgain, err := time.Parse(time.RFC3339, nowString)
	fmt.Println(err)
	fmt.Println(nowAgain)
	fmt.Println(nowAgain.Format(time.Kitchen))

	date1 := "2020-04-19T18:24:09+02:00"
	date2 := "2020-04-18T18:24:09+02:00"
	date3 := "2021-04-18T18:24:09+02:00"

	timeDate1, _ := time.Parse(time.RFC3339, date1)
	timeDate2, _ := time.Parse(time.RFC3339, date2)
	timeDate3, _ := time.Parse(time.RFC3339, date3)

	var dates []time.Time

	myDates := append(dates, timeDate1, timeDate2, timeDate3)

	fmt.Println(myDates)

	sort.Slice(myDates, func(i, j int) bool {
		return myDates[i].Before(myDates[j])
	})

	fmt.Println(myDates)

}
