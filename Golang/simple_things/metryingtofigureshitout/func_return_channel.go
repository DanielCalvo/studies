package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Img struct {
	Img      image.Image
	Filename string
	SrcDir   string
	DstDir   string
	Ratio    int
}

//TODO: Do error handling properly (print out more useful information)

func main() {

	fmt.Println("hi")
	ImgGen("/tmp/img", "/tmp/resized/", 20, 50)

}

func ImgGen(srcDir, dstDir string, ratios ...int) <-chan Img {
	//func ImgGen(srcDir, dstDir string, ratios ...int) {

	if !strings.HasSuffix(srcDir, "/") {
		srcDir = srcDir + "/"
	}
	if !strings.HasSuffix(dstDir, "/") {
		dstDir = dstDir + "/"
	}

	_, err := os.Stat(srcDir)
	if os.IsNotExist(err) {
		log.Panic(err)
	}

	_, err = os.Stat(dstDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dstDir, 0755)
		if err != nil {
			log.Panic(err)
		}
	}

	for _, rt := range ratios {
		_, err := os.Stat(dstDir + strconv.Itoa(rt))
		if os.IsNotExist(err) {
			err = os.Mkdir(dstDir+strconv.Itoa(rt), 0755)
			if err != nil {
				log.Panic(err)
			}
		}
	}

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	out := make(chan Img)
	//go func here
	go func() {
		for _, file := range files {
			for _, rt := range ratios {
				_, err := os.Stat(dstDir + strconv.Itoa(rt) + "/" + file.Name())
				if os.IsNotExist(err) {
					fmt.Println("Doesn't exist, would resize:", file.Name())

					img, err := imgio.Open(srcDir + file.Name())
					if err != nil {
						continue
					}

					imgc := Img{
						Img:      img,
						Filename: file.Name(),
						SrcDir:   srcDir,
						DstDir:   dstDir,
						Ratio:    rt,
					}
					out <- imgc
				}
			}
		}
		close(out)
	}()
	return out
}

//
//
//	if !file.IsDir() && strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".jpeg") {
//		//put on channel
//	}
//}
//
//
//
//go func() {
//	img := Img {
//		Img: nil, //TODO
//		Filename: "TODO",
//		SrcDir: srcDir,
//		DstDir: dstDir,
//		Ratios: ratios,
//	}
//
//	out <- img
//	close(out)
//}()
//
//
//return out

//c1 := gen(1,2,3)
//for c2 := range c1 {
//fmt.Println(c2)
//}
