package image_common

import (
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"image"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Img struct {
	Img      image.Image
	Filename string
	SrcDir   string
	DstDir   string
	Ratio    int
}

func ImgGen(srcDir, dstDir string, ratios []int) <-chan Img {

	if !strings.HasSuffix(srcDir, string(os.PathSeparator)) {
		srcDir = srcDir + string(os.PathSeparator)
	}
	if !strings.HasSuffix(dstDir, string(os.PathSeparator)) {
		dstDir = dstDir + string(os.PathSeparator)
	}

	_, err := os.Stat(srcDir)
	if os.IsNotExist(err) {
		log.Fatalf("Directory with images to be resized (%s) does not exist\n", srcDir)

	}

	_, err = os.Stat(dstDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dstDir, 0755)
		if err != nil {
			log.Fatalf("Cannot create destination directory at %s", dstDir)
		}
	}

	for _, rt := range ratios {
		_, err := os.Stat(dstDir + strconv.Itoa(rt))
		if os.IsNotExist(err) {
			err = os.Mkdir(dstDir+strconv.Itoa(rt), 0755)
			if err != nil {
				log.Fatalf("Cannot create destination directory at %s", dstDir+strconv.Itoa(rt))
			}
		}
	}

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatalf("Unable to read the files from %s", srcDir)
	}

	out := make(chan Img)
	go func() {
		for _, file := range files {
			for _, rt := range ratios {
				_, err := os.Stat(dstDir + strconv.Itoa(rt) + string(os.PathSeparator) + file.Name())
				if os.IsNotExist(err) {
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
				} else {
					log.Println("Already resized:", file.Name())
				}
			}
		}
		close(out)
	}()
	return out
}

func ImgRes(c <-chan Img, wg *sync.WaitGroup) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer wg.Done()
			for i := range c {
				log.Println("Resizing:", i.Filename)
				iRes := transform.Resize(i.Img, i.Img.Bounds().Dx()*i.Ratio/100, i.Img.Bounds().Dy()*i.Ratio/100, transform.Linear)
				err := imgio.Save(i.DstDir+strconv.Itoa(i.Ratio)+string(os.PathSeparator)+i.Filename, iRes, imgio.JPEGEncoder(82))
				if err != nil {
					log.Fatalf("Unable to save file. Error: %s", err)
				}
			}
		}()
	}
}
