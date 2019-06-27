package main

import (
	"encoding/json"
	"fmt"
)

//http://api.open-notify.org/astros.json

type astronauts struct {
	People []struct {
		Craft string `json:"craft"`
		Name  string `json:"name"`
	} `json:"people"`
	Number  int    `json:"number"`
	Message string `json:"message"`
}

func main() {

	myjson := `{"people": [{"craft": "ISS", "name": "Alexey Ovchinin"}, {"craft": "ISS", "name": "Nick Hague"}, {"craft": "ISS", "name": "Christina Koch"}], "number": 3, "message": "success"}`
	myjson_bytes := []byte(myjson)
	astronauts1 := astronauts{}

	json.Unmarshal(myjson_bytes, &astronauts1)

	fmt.Println("Did it work?")
	fmt.Println(astronauts1)

	for _, v := range astronauts1.People {
		fmt.Println(v)
	}

}
