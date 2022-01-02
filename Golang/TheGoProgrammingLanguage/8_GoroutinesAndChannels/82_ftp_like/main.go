// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 219.
//!+

// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//To connect: nc localhost 8000
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		if input.Text() == "ls" {
			//Get files in local dir, send list back
			fmt.Fprint(c, "You sent LS!"+"\n")
		}

		if strings.HasPrefix("cat", input.Text()) {
			fmt.Fprintln(c, "Implementation!")
		}

		fmt.Fprint(c, input.Text()+"\n")
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
