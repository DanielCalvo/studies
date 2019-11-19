package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
)
import "google.golang.org/grpc"
import "../image_resizerpb"

func main() {
	fmt.Println("Hi from client!")

	file, err := ioutil.ReadFile("/tmp/a.jpeg")

	req := &image_resizerpb.ImgRequest{
		ImgName: "Asd",
		Img:     file,
	}
	fmt.Println("Sending as request:", req.GetImgName())

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	c := image_resizerpb.NewImgServiceClient(cc)
	res, err := c.Resize(context.Background(), req)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Got response:", res.ImgName)

}
