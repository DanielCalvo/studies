package main

import (
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	origDir := "/tmp/img"
	resizedDir := origDir + "/resized"

	_, err := os.Stat(resizedDir)
	if os.IsNotExist(err) {
		os.Mkdir(resizedDir, 0775)
	}

	files, err := ioutil.ReadDir(origDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		originalFilePath := origDir + "/" + file.Name()
		resizedFilePath := resizedDir + "/" + file.Name()

		if !strings.HasSuffix(file.Name(), ".jpg") {
			log.Println(file.Name(), "does not end with .jpg, skipping")
			continue
		}

		if _, err := os.Stat(resizedFilePath); err == nil {
			log.Println(file.Name(), "is already resized, skipping")
			continue
		}

		img, err := imgio.Open(originalFilePath)
		if err != nil {
			log.Println("Could not open", originalFilePath)
			log.Println("Error:", err)
			continue
		}

		if img.Bounds().Dx()+img.Bounds().Dy() == 7084 {
			imgResized := transform.Resize(img, img.Bounds().Dx()/2, img.Bounds().Dy()/2, transform.Linear)
			e := imgio.Save(resizedFilePath, imgResized, imgio.JPEGEncoder(82))
			log.Print("Resized ", resizedFilePath)
			if e != nil {
				log.Println("There was en error resizing:", resizedFilePath)
				log.Println("Error:", err)
			}
		}
	}
}
