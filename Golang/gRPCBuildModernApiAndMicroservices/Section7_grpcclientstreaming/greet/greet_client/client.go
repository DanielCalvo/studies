package main

import (
	"../greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	fmt.Println("Hi from client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Dani",
				LastName:  "McDani",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Not Dani",
				LastName:  "Not McDani",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Bananas",
				LastName:  "And apples!",
			},
		},
	}

	for _, req := range requests {
		fmt.Println("Sending req:", req)
		err = stream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("LongGreet response:", res.Result)

}
