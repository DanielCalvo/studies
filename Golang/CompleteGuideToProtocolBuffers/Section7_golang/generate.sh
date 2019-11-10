#!/usr/bin/env bash
#Run from current dir!
protoc -I ./ --go_out=./simple ./simple/simple.proto
