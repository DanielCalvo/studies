package main

import (
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

func main() {
	img, err := imgio.Open("/home/daniel/Downloads/imeu09mf88n21.jpg")
	if err != nil {
		panic(err)
	}

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	img_resized := transform.Resize(img, width/2, height/2, transform.Linear)

	if e := imgio.Save("/tmp/img_resized.jpg", img_resized, imgio.JPEGEncoder(82)); e != nil {
		panic(err)
	}
}
