package main

import (
	"example/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

/*
curl -X POST -H "Content-Type: application/json" -d '{"name":"Bob McBobson","gender":"male","age":11}' http://localhost:8080/user
curl http://localhost:8080/user/62299ff8fd09d115cfcaa828
curl -X DELETE http://localhost:8080/user/62299ff8fd09d115cfcaa828


*/

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
