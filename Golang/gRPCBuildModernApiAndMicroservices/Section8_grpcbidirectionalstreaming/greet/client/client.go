package main

import (
	"../greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("You called client!")

	fmt.Println("Hi from client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Dani",
				LastName:  "McDani",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Not Dani",
				LastName:  "Not McDani",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Bananas",
				LastName:  "And apples!",
			},
		},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			fmt.Println("Sending message with FirstName:", req.GetGreeting().GetFirstName())
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err != io.EOF {
				close(waitc)
			}
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("Received: %v", res.GetResult())
		}
	}()

	<-waitc

}
