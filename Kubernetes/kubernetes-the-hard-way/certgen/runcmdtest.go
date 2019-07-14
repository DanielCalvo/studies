package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	mycmd := exec.Command("pwd")

	bs, err := mycmd.CombinedOutput()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))

}
