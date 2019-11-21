package main

import (
	"../greetpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct{}

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

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("You called GreetEveryone!")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName
		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalln(err)
			return err
		}

	}

	return nil
}
