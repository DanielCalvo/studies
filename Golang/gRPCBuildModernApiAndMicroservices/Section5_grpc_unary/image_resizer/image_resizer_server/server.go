package main

import (
	"../image_resizerpb"
	"bytes"
	"context"
	"fmt"
	"github.com/anthonynsimon/bild/transform"
	"google.golang.org/grpc"
	"image"
	"image/jpeg"
	"log"
	"net"
)

type server struct{}

func (*server) Resize(ctx context.Context, req *image_resizerpb.ImgRequest) (*image_resizerpb.ImgResponse, error) {
	fmt.Println("Received:", req.GetImgName())

	r := bytes.NewReader(req.Img)
	img, _, err := image.Decode(r)
	if err != nil {
		fmt.Println("Could not decode image:", req.GetImgName())
	}

	imgResized := transform.Resize(img, img.Bounds().Dx()*50/100, img.Bounds().Dy()*50/100, transform.Linear)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, imgResized, nil)
	if err != nil {
		log.Fatalln(err)
	}

	res := &image_resizerpb.ImgResponse{
		ImgName: req.GetImgName(),
		Img:     buf.Bytes(),
	}

	return res, nil
}

func main() {
	fmt.Println("Hello from image server!")

	s := grpc.NewServer()
	image_resizerpb.RegisterImgServiceServer(s, &server{})

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("End of main for server reached")
}
