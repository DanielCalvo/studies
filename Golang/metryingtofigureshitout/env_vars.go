package main

import (
	"fmt"
	"os"
)

func main() {

	envVars := []string{"GOPATH", "GOROOT", "PWD", "myvar", "somevar", "someip"} // so this works if I run it on my local(TM), but inside a docker container I don't get the variables. :(

	for i, v := range envVars {
		if os.Getenv(v) != "" {
			fmt.Println("Found:", i, v, os.Getenv(v))
		} else {
			fmt.Println("Not found:", i, v)
		}
	}
}
