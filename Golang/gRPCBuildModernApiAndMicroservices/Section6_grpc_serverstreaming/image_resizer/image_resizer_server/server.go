package main

import (
	"../image_resizerpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

//func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
//
//}

func (*server) Resize(req *image_resizerpb.ImgRequest, stream image_resizerpb.ImgService_ResizeServer) error {
	//magic goes here
	return nil
}

func main() {
	fmt.Println("Hello!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	image_resizerpb.RegisterImgServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("End of main for server reached")

}
