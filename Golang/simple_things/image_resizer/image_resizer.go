package main

import (
	"./image_common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"
)

type Parameters struct {
	SrcDir string `yaml:"SrcDir"`
	DstDir string `yaml:"DstDir"`
	Ratio  []int  `yaml:"Ratio"`
}

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Could not get current directory")
	}

	parametersFile, err := ioutil.ReadFile(pwd + string(os.PathSeparator) + "parameters.txt")
	if err != nil {
		log.Fatalln("Could not find parameters file\nLooked for it at: " + pwd + string(os.PathSeparator) + "parameters.txt")
	}

	var parameters Parameters
	err = yaml.Unmarshal(parametersFile, &parameters)
	if err != nil {
		log.Fatalln("parameters.txt has an invalid format")
	}

	c := image_common.ImgGen(parameters.SrcDir, parameters.DstDir, parameters.Ratio)
	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())
	image_common.ImgRes(c, &wg)
	wg.Wait()

}
