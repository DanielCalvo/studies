package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go myhandle(conn)
	}
}

func myhandle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			fmt.Println("Got GET")
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET":
			fmt.Println("Got SET")
			if len(fs) != 3 {
				fmt.Fprintf(conn, "EXPECTED VALUE\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			fmt.Println("Got DEL")
		case "exit":
			conn.Close()
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
