package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you!") //Hey you could use this to send a json object and see what gives!
	//That, or curl it yourself
}
