package main

import (
	"fmt"
	"net/http"
)

type hotdog int

//http.ListenAndServe takes anything of type handler
//Type handler interface: Any type that has a method with this signature: ServerHTTP(ResponseWriter, *Request)

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want can go in this function")

}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
