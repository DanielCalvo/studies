package main

import (
	"../greetpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("Long greet function was invoked!")
	result := "Hello "
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			//We have finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalln(err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += firstName + "! "

	}

}

func main() {
	fmt.Println("Hello!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("End of main for server reached")

}
