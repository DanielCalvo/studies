package main

import (
	"log"
	"os/exec"
)

func GitClone(repo, filepath string) error {

	//git clone https://github.com/kubernetes/kubectl banana/kubectl_custom

	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)

	return nil
}

func main() {

	//err := GitClone("https://github.com/kubernetes/kubectl", "/tmp/testing")
	//
	//if err != nil {
	//	fmt.Println("Error:",err)
	//	os.
	//}

}
