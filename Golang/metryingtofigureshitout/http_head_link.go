package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	link := flag.String("link", "", "HTTP link to check with http.Head")
	flag.Parse()

	resp, err := http.Head(*link)
	if err != nil {
		fmt.Println("---- GOT ERROR -----")
		fmt.Println("Error:", err)
	}

	fmt.Println(resp.Status)
	//fmt.Println(resp.Body)

}
