package main

import (
	"flag"
	"fmt"
)

func main() {

	//go run webserver.go -webserver

	port := flag.Int("webserver-port", 0, "Port for the webserver to listen on")
	repo := flag.String("repo", "", "Individual repository to scan markdown files and generate results")

	flag.Parse()

	if *port != 0 {
		fmt.Println("Webserver flag was set! Running webserver!")
		//run webserver
	}

	if *repo != "" {
		fmt.Println("Individual repo flag was set. Scanning repo:", *repo)
		//scan single repo
	}

	fmt.Println("No flags set!")
	fmt.Println("Running defaults: Scanning all repositories configured in repositories.yaml")

}
