package main

import (
	"log"
	"net/http"
)

func main() {
	//index_old.html is treated differently
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
