package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	//Marshalling a struct
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	//Let's unmarshall it!
	var m1 Message
	err = json.Unmarshal(b, &m1)
	if err != nil {
		panic(err)
	}
	fmt.Println(m1)

	//But what if you need to unmarshall some json that you don't know the types it has?
	b1 := []byte(`{"Name":"Wednesday","Age":6,"Parents":["","Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b1, &f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b1))

	m2 := f.(map[string]interface{})
	fmt.Println(m2)

	for k, v := range m2 {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
