package main

import (
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println(r)
	//r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)

}
