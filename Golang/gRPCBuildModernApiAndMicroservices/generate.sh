#!/usr/bin/env bash

#Needs to be run from the same directory as the script (${PWD}/generate.sh)
protoc Section5_grpc_unary/greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc Section5_grpc_unary/calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.

protoc Section5_grpc_unary/image_resizer/image_resizerpb/image_resizer.proto --go_out=plugins=grpc:.

protoc Section6_grpc_serverstreaming/greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc Section6_grpc_serverstreaming/image_resizer/image_resizerpb/image_resizer.proto --go_out=plugins=grpc:.

protoc Section7_grpcclientstreaming/greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc Section8_grpcbidirectionalstreaming/greet/greetpb/greet.proto --go_out=plugins=grpc:.
