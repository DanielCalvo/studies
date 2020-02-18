package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You have reached Dani's webserver!!!!!")

}

func main() {
	fmt.Println("Starting webserver!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)

}
