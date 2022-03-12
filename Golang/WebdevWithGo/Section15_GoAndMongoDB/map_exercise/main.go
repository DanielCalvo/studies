package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type User map[int]struct {
	Name string
	Age  int
}

var user = make(User)
var i int
var m sync.Mutex

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", GetUser)       //http://localhost:8080/user/0
	r.POST("/user", CreateUser)       //curl -X POST -H "Content-Type: application/json" -d '{"Name":"Bob McBobson","Age":11}' http://localhost:8080/user
	r.DELETE("/user/:id", DeleteUser) //curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777
	log.Fatalln(http.ListenAndServe("localhost:8080", r))

}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	i, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}

	mj, err := json.Marshal(user[i])
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

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	httpId := p.ByName("id")

	id, err := strconv.Atoi(httpId)
	if err != nil {
		log.Println("Unable to convert user id to int:", httpId)
		return
	}

	if _, ok := user[id]; !ok {
		log.Println("User ID does not exist:", id)
		fmt.Fprint(w, "User ID does not exist: ", id, "\n")
		w.WriteHeader(404)
		return
	}

	delete(user, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user with id ", httpId, "\n")
}
