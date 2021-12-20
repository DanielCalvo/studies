package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	postBody, _ := json.Marshal(map[string]string{
		"Name":     "Daniel",
		"Interest": "Advances Memes",
	})

	bodyBuf := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://postman-echo.com/post", "application/json", bodyBuf)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
