package image_common

import (
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/pkg/errors"
	"image"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func SayHi(s string) {
	fmt.Println("This is a sample function that just says hi!")
}

type Img struct {
	Img   image.Image
	Ratio int
}

func ResizeImage(img image.Image, ratios ...int) []Img {
	//fmt.Println("Hi from ResizeImage")
	var imgs []Img
	for _, rt := range ratios {
		imgRes := transform.Resize(img, img.Bounds().Dx()*rt/100, img.Bounds().Dy()*rt/100, transform.Linear)

		i := Img{
			Img:   imgRes,
			Ratio: rt,
		}
		imgs = append(imgs, i)
	}
	return imgs
}

//A comment on how concurrency and file saving are handled would be nice here!
func ResizeDir(dir string, ratios ...int) error {

	if len(ratios) == 0 {
		return errors.New("ratios parameter must be at least one integer")
	}

	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}

	for _, rDir := range ratios {
		_, err := os.Stat(dir + strconv.Itoa(rDir))
		if os.IsNotExist(err) {
			err = os.Mkdir(dir+strconv.Itoa(rDir), 0755)
			if err != nil {
				return err
			}
		}
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {

		img, err := imgio.Open(dir + file.Name())
		if err != nil {
			continue
		}

		//multi-routine logic would go here

		Imgs := ResizeImage(img, ratios...)

		fmt.Println("Resizing:", file.Name(), "with ratios:", ratios, "goroutines:", runtime.NumGoroutine())

		for _, i := range Imgs {
			err = imgio.Save(dir+strconv.Itoa(i.Ratio)+"/"+file.Name(), i.Img, imgio.JPEGEncoder(82))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
