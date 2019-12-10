package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

//Can you create a struct with image sizes and the resize ratio that you want for them?

func main() {

	srcDir := "/media/daniel/ruhroh/Daniel O RETORNO/Imagens/mobile pix/"
	dstDir := "/media/daniel/ruhroh/Daniel O RETORNO/Imagens/mobile pix/resized/"
	resizeRatioPercent := 50

	_, srcDirErr := os.Stat(srcDir)
	if srcDirErr != nil {
		log.Fatal("Error accessing source directory: ", srcDirErr)
	}

	_, dstDirErr := os.Stat(dstDir)
	if os.IsNotExist(dstDirErr) {
		dstDirMkdirErr := os.Mkdir(dstDir, 0755)
		if dstDirMkdirErr != nil {
			log.Fatal("Could not create destination directory:", dstDirMkdirErr)
		}
		fmt.Println("Created destination directory: ", dstDir)
	} else if dstDirErr != nil {
		log.Fatal("Error with destination directory. Error:", dstDirErr)
	}

	imageChan := make(chan string)

	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())

	go processImages(imageChan, &wg, resizeRatioPercent, srcDir, dstDir)
	getImages(imageChan, srcDir, dstDir)

	wg.Wait()

	fmt.Println("End of main")

}

func getImages(c chan string, srcDir, dstDir string) {
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if _, err := os.Stat(dstDir + file.Name()); err == nil {
			continue
		}
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".jpeg") {
			c <- file.Name()
		}
	}
	close(c)
}

func processImages(imageChan chan string, wg *sync.WaitGroup, resizeRatioPercent int, srcDir, dstDir string) {
	//Improve print messages to have better formatting

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer wg.Done()
			for v := range imageChan {
				resizeImage(v, srcDir, dstDir, resizeRatioPercent)
			}
		}()
	}
}

func resizeImage(imgFileName, srcDir, dstDir string, resizeRatioPercent int) {
	img, err := imgio.Open(srcDir + imgFileName)
	if err != nil {
		fmt.Println("Could not open", srcDir+imgFileName)
		fmt.Println("Error:", err)
	}

	imgResized := transform.Resize(img, img.Bounds().Dx()*resizeRatioPercent/100, img.Bounds().Dy()*resizeRatioPercent/100, transform.Linear)
	e := imgio.Save(dstDir+imgFileName, imgResized, imgio.JPEGEncoder(82))
	if e != nil {
		fmt.Println("There was en error resizing:", dstDir+imgFileName)
		fmt.Println("Error:", err)
	}

	fmt.Println("Resized:", imgFileName, runtime.NumGoroutine())
}
