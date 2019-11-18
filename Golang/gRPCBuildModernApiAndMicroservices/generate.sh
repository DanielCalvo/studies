#!/usr/bin/env bash

#Needs to be run from the same directory as the script (${PWD}/generate.sh)
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.


protoc image_resizer/image_resizerpb/image_resizer.proto --go_out=plugins=grpc:.