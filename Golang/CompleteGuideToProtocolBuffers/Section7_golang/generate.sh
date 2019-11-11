#!/usr/bin/env bash
#Run from current dir!
protoc -I ./ --go_out=./ ./simple/simple.proto
protoc -I ./ --go_out=./ ./enum_example/enum_example.proto
protoc -I ./ --go_out=./ ./complex/complex.proto
