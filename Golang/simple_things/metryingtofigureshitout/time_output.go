package main

import (
	"fmt"
	"time"
)

func main() {
	myTime := time.Now()
	fmt.Println(myTime.Format(time.RFC3339))
	fmt.Println(time.Now().Format(time.RFC3339))
}

//2020-04-12T14:28:01+02:00
//2020-04-12 00:12:59.044259689 +0000 UTC
