package main

import (
	"../image_resizerpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Println("Hi from client!")

	imgDir := "/tmp/img/"

	files, err := ioutil.ReadDir(imgDir)
	if err != nil {
		log.Fatalln(err)
	}

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	c := image_resizerpb.NewImgServiceClient(cc)

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".jpeg") {

			imgContent, err := ioutil.ReadFile(imgDir + file.Name())
			if err != nil {
				fmt.Println("Could not read:", file.Name())
				continue
			}

			req := &image_resizerpb.ImgRequest{
				ImgName: file.Name(),
				Img:     imgContent,
			}

			res, err := c.Resize(context.Background(), req)

			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("Got response:", res.GetImgName())

			err = ioutil.WriteFile(imgDir+"res_"+file.Name(), res.Img, 0644)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

}
