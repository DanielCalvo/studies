package main

import (
	"../greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
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

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dani",
			LastName:  "DaniDani",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Response from greet: %v", res.Result)

	req2 := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dani",
			LastName:  "McDani",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req2)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//We've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		fmt.Println("Response from GreetManyTimes:", msg.GetResult())
	}

}
