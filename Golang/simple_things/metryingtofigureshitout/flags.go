package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	hosts     string
	namespace string
	port      int
	service   string
)

var banana *string

func init() {
	flag.StringVar(&hosts, "hosts", "giantswarm.io,kubernetes.default.svc.cluser.local", "DNS hosts to resolve")
	flag.StringVar(&namespace, "namespace", "monitoring", "Namespace of net-exporter service")
	flag.IntVar(&port, "port", 8000, "Port of net-exporter service")
	flag.StringVar(&service, "service", "net-exporter", "Name of net-exporter service")
	banana = flag.String("banana", "bananaaaaa", "just banana things")
	fmt.Printf("%T", banana)
	fmt.Printf(*banana)
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "--help") {
		fmt.Println("No bueno")
		return
	}
	flag.Parse()

	fmt.Println("Hosts: ", hosts)
	fmt.Println("Namespace: ", namespace)
	fmt.Println("Port: ", port)
	fmt.Println("Service: ", service)
	fmt.Println("This line contains banana: ", *banana)

	for _, c := range os.Args {
		fmt.Println("Printing type result for: ", c)
		fmt.Printf("%T", c)
		fmt.Println()
	}

}
