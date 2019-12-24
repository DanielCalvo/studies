package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

//http://markdownscanner/project/repo

//maybe have a list of repos you want to go through on an init function?
//cache the results and if the results are older than x minutes, re-do the request?

func init() {
	log.Println("memes")
}

func main() {
	//depending on the path you get, show a file
	//if a file doesn't exist, run a function to generate the file

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))

}

//map a URL path to a given file/report?

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
	log.Println(req.URL)

	if req.URL.Path == "/kubernetes/kubectl" {

		//if report already exists
		//show it

		//if report does not exist
		//run it
		//respond with report?

		log.Println("yeah")
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, "anything you want")

	}
}
