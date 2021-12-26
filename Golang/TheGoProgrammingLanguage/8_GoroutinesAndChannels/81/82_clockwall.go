package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must provide at least one argument")
		os.Exit(1)
	}

	for _, v := range os.Args[1:] {
		go func() {
			conn, err := net.Dial("tcp", v)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			mustCopy(os.Stdout, conn)
		}()
	}
	//Sleeping here as we have no channel to block :(
	time.Sleep(time.Minute)

}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
