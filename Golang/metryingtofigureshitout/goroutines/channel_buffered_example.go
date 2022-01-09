package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	responses := make(chan string, 3)
	go func() { responses <- request("https://google.com") }()
	go func() { responses <- request("https://facebook.com") }()
	go func() { responses <- request("https://apple.com") }()

	fmt.Println(<-responses) //Prints whoever responded first, cool!
}

func request(s string) string {
	_, err := http.Get(s)
	if err != nil {
		time.Sleep(time.Second)
		return "Site " + s + " errored: " + err.Error()
	}
	//You could aso:
	//buf := bytes.Buffer{}
	//buf.ReadFrom(resp.Body) //ignoring error
	//return buf.String()
	return s
}
