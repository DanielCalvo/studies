package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//this comes from the go std library example here: https://pkg.go.dev/net/http#example-Get
	//but with the example get as it, you cant set headers, so you need to create the request, then set headers, then do the call
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
