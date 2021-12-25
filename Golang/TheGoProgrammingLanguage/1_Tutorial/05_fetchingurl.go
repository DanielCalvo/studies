package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//go run 05_fetchingurl.go https://google.com
//Started with the code snippet in the book but then modified it to do the exercises
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get("https://" + url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//_, err = ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err to stdout: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Status code:", resp.Status)

	}
}
