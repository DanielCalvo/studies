package main

import (
	"../calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type watevs struct{}

//func (*watevs) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
//
//}

func (*watevs) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	//Logic goes here:
	mySum := req.GetFirstNumber() + req.GetSecondNumber()

	res := &calculatorpb.SumResponse{
		Result: mySum,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &watevs{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("End of main for server reached")
}
