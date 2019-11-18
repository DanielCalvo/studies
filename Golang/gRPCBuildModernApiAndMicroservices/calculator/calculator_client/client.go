package main

import (
	"../calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hi from client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	fmt.Printf("Created client: %f", c)

	req := &calculatorpb.SumRequest{
		FirstNumber:  4,
		SecondNumber: 2,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Response from greet: %v\n", res.Result)

	fmt.Println("Result:", res.Result)

}
