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

}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":80", nil))
}
