package main

import (
	"log"
	"os"
	"text/template"
)

type server struct {
	Hostname string
	IP       string
}

func main() {
	tpl, err := template.ParseFiles("/home/daniel/PycharmProjects/studies/Golang/WebdevWithGo/Section2_templates/11_datastructuresintotemplates/myservers.txt")
	if err != nil {
		log.Fatal(err)
	}

	localhost := server{
		Hostname: "localhost",
		IP:       "127.0.0.1",
	}
	router := server{
		Hostname: "router",
		IP:       "192.168.1.1",
	}

	servers := []server{localhost, router}

	err = tpl.Execute(os.Stdout, servers)
	if err != nil {
		log.Fatal(err)
	}
}
