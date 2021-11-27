package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//Let's create and import a package
//And then get user in a kinda dummy way, the way todd did it!

func main() {
	r := httprouter.New()
	r.GET("/", index)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")

}
