package main

import (
	"fmt"
	"net/http"
)

//http://markdownscanner/project/repo

//maybe have a list of repos you want to go through on an init function?
//cache the results and if the results are older than x minutes, re-do the request?
//depending on the path you get, show a file
//if a file doesn't exist, run a function to generate the file
//map a URL path to a given file/report?

//https://stackoverflow.com/questions/46516797/additional-arguments-to-http-function-golang
//https://stackoverflow.com/questions/26211954/how-do-i-pass-arguments-to-my-handler

//rename those fields pls
//get config from yaml and/or env variables

//put all the config on a single struct and later see what gives

//type Config struct {
//	TmpDir string
//}
//
//var config = Config {
//	TmpDir: "/tmp/mdscanner",
//}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Would've opened:", config.tmpDir, req.URL.Path, ".json")
}

///tmp/mdscanner/results/kubernetes/kubeadm/report.json

func main() {

	http.HandleFunc("/", handler)
	//http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

//func foo(w http.ResponseWriter, req *http.Request) {
//	log.Println(req.URL)
//
//	if req.URL.Path == "/kubernetes/kubectl" {
//
//		dat, err := ioutil.ReadFile("/tmp/kubectl.json")
//		if err != nil {
//			fmt.Fprintln(w, err)
//			return
//		}
//		fmt.Fprintln(w, string(dat))
//	}
//}
