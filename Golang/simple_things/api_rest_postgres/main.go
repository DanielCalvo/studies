package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type server struct{}

type Company struct {
	Name       string
	Valuation  float64
	DateJoined string
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:example@localhost/startups?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connection to database established!")
}

func main() {
	//s := &server{}
	//http.Handle("/", s)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	//...
	api.HandleFunc("/companies/name/{name}", searchByName).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe(":8080", r))

}

var id int
var name, valuation, datejoined string

//http://localhost:8080/api/v1/companies/name/SpaceX
//select * from company where name = 'SpaceX';
func searchByName(w http.ResponseWriter, r *http.Request) {
	var c Company

	pathParams := mux.Vars(r)
	if name, ok := pathParams["name"]; ok {
		fmt.Println("Got name:", name)
		row := db.QueryRow("SELECT * FROM company WHERE name = $1", name)

		err := row.Scan(&id, &c.Name, &c.Valuation, &c.DateJoined)
		switch {
		case err == sql.ErrNoRows:
			http.NotFound(w, r)
			return
		case err != nil:
			fmt.Println("Error querying DB:", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(c)
		if err != nil {
			fmt.Println("Error marshalling json:", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}

		fmt.Fprintln(w, string(b))
	}
}

//func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte(`{"message": "hello world"}`))
//}
