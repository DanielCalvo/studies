package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "api v1")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
