package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

//Uhhh this is called certgen but in the end I'm doing everything here...

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

// Checks if a given command outputs a string. If it does, we return true. If it does not, we return false.
func validateCmd(successString string, commandArguments ...string) bool {
	fmt.Println("successString", successString)
	fmt.Println("Actual command:", commandArguments[0])
	fmt.Println("Command parameters", commandArguments[1:])

	cmd := exec.Command(commandArguments[0], commandArguments[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return false
	}
	r := bytes.NewReader(stdoutStderr)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), successString) {
			fmt.Println("Worked!")
			return true
		}
	}
	return false
}

func main() {

	//gcloud commands to run on your local before all of this
	//Still to automate the gcloud stuff:
	/*

		RUN ONCE:
		gcloud init

		NETWORK SET UP:
		gcloud compute networks create kubernetes-the-hard-way --subnet-mode custom
		gcloud compute networks subnets create kubernetes \
		  --network kubernetes-the-hard-way \
		  --range 10.240.0.0/24
		gcloud compute firewall-rules create kubernetes-the-hard-way-allow-internal \
		  --allow tcp,udp,icmp \
		  --network kubernetes-the-hard-way \
		  --source-ranges 10.240.0.0/24,10.200.0.0/16
		gcloud compute firewall-rules create kubernetes-the-hard-way-allow-external \
		  --allow tcp:22,tcp:6443,icmp \
		  --network kubernetes-the-hard-way \
		  --source-ranges 0.0.0.0/0
		gcloud compute addresses create kubernetes-the-hard-way \
		  --region $(gcloud config get-value compute/region)
		gcloud compute addresses list --filter="name=('kubernetes-the-hard-way')"

	*/

	//gcloud_cmd := `echo`

	//func validateCmd (successString string, commandArguments ...string) bool {
	validateCmd("banana", "echo", "banana", "apple", "grape")

	//h := Parameters{}
	//
	//yamlFile, err := ioutil.ReadFile("/home/daniel/PycharmProjects/studies/Kubernetes/kubernetes-the-hard-way/certgen/parameters.yaml")
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//
	//err = yaml.Unmarshal(yamlFile, &h)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//
	//fmt.Println(h.Hosts)
	//
	//for _, v := range h.Hosts.Workers {
	//	fmt.Println(v.Hostname)
	//}

	//Next actions here:
	//Find a way to properly organize the way you run the commands to generate the certificates
	//Define what meaningful values you want on the .yaml file
	//Create the certificates with the values you got from the .yaml file.

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
