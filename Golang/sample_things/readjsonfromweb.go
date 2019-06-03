package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json:"number"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"
	spaceClient := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	//Body is a slice of bytes! This means I should be able to write it to a file!
	fileErr := ioutil.WriteFile("/home/daniel/PycharmProjects/studies/Golang/sample_things/myjson.json", body, 0644)
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(people1.Number)

}
