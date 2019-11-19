package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
)
import "../image_resizerpb"

type server struct {
}

func (*server) Resize(ctx context.Context, req *image_resizerpb.ImgRequest) (*image_resizerpb.ImgResponse, error) {
	fmt.Println("Received:", req.GetImgName())

	err := ioutil.WriteFile("/tmp/banana.jpeg", req.Img, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	res := &image_resizerpb.ImgResponse{
		ImgName: "Banana",
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
