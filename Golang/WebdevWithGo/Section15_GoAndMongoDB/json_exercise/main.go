package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type User map[int]struct {
	Name string
	Age  int
}

var user = make(User)
var m sync.Mutex
var i int

func main() {

	/*
		When the program starts, read the file if it exists and unmarshall the json into memory
		When you run GetUser, you just read the map and send the result
		When you run CreateUser, update the map and save it to the json file
		When you run DeleteUser, update the map and then save it to the json file

		Do you read again from the JSON file at any point? I think no right?

	*/

	err := ReadJsonFromFile(&user, "test.json")
	if err != nil {
		log.Println("Unable to read json from file:", err)
	}

	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", GetUser) //http://localhost:8080/user/0
	//r.POST("/user", CreateUser) //curl -X POST -H "Content-Type: application/json" -d '{"Name":"Bob McBobson","Age":11}' http://localhost:8080/user
	//r.DELETE("/user/:id", DeleteUser) //curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}

func ReadJsonFromFile(u *User, fsPath string) error {
	jsonFile, err := os.Open(fsPath)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &u)
	if err != nil {
		return err
	}
	return nil
}

func WriteJsonToFile(u *User, fsPath string) error {
	b, err := json.MarshalIndent(*u, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fsPath, b, 0644) //Uh, truncates all the time, not what you want, you want append!
	if err != nil {
		return err
	}

	return nil
}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(user) //Empty, user is being changed in main but not in here, even though it's a global variable, ffs
	httpId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}

	mj, err := json.Marshal(user[httpId])
	fmt.Println(user[httpId])
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", mj)
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read from body:", err)
		return
	}

	u := struct {
		Name string
		Age  int
	}{}

	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Println("Unable to unmarshall json:", err)
		return
	}

	user[i] = u
	m.Lock()
	i++
	m.Unlock()

	fmt.Fprintf(w, "%+v\n", u)
	w.Header().Set("Content-Type", "application/json")
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}
