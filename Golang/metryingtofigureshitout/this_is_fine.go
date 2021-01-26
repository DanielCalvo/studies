package main

import (
	"bytes"
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

func fine(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "this_is_fine.jpg")
	log.Println("this_is_fine.jpg")

	img, err := imgio.Open("/home/daniel/PycharmProjects/studies/Golang/simple_things/metryingtofigureshitout/this_is_fine.jpg")
	if err != nil {
		log.Println("Could not open", "/home/daniel/PycharmProjects/studies/Golang/simple_things/metryingtofigureshitout/this_is_fine.jpg")
		log.Println("Error:", err)
	}

	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, img, nil)

	if err != nil {
		fmt.Println("Unable to encode image")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	w.Write(buffer.Bytes())
}

func main() {
	http.HandleFunc("/fine", fine)
	http.ListenAndServe(":8080", nil)

}
