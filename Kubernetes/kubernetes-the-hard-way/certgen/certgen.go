package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Parameters struct {
	Hosts struct {
		Controllers []struct {
			Hostname string `yaml:"hostname"`
			IP       string `yaml:"ip"`
		} `yaml:"controllers"`
		Workers []struct {
			Hostname string `yaml:"hostname"`
			IP       string `yaml:"ip"`
		} `yaml:"workers"`
	} `yaml:"hosts"`
}

func main() {

	h := Parameters{}

	yamlFile, err := ioutil.ReadFile("/home/daniel/PycharmProjects/studies/Kubernetes/kubernetes-the-hard-way/certgen/parameters.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &h)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(h.Hosts)

	for _, v := range h.Hosts.Workers {
		fmt.Println(v.Hostname)
	}

	//emptyFile, err := os.Create("/tmp/myfile.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err1 := emptyFile.Close()
	//if err1 != nil {
	//	log.Fatal(err1)
	//}

	//Create file on somedir
}
