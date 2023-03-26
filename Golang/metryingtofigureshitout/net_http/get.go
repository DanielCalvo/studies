package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error ", err)
	}

	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body) //ignoring error
	fmt.Println(buf.String())

}
