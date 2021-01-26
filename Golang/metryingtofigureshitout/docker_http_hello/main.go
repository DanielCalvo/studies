package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//Cool article: https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")

	hostName, err := os.Hostname()
	if err != nil {
		hostName = "Hostname could not be retrieved"
	}
	fmt.Fprintln(w, "Hostname: "+hostName)

	t := time.Now()
	fmt.Fprintln(w, "Current time: "+t.Format("20060102150405"))

	fmt.Fprintln(w, "name:", os.Getenv("name"))
	fmt.Fprintln(w, "surname:", os.Getenv("surname"))
	fmt.Fprintln(w, "age:", os.Getenv("age"))

}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
