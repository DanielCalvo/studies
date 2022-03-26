package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type server struct{}

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

//http://localhost:8080/api/v1/companies/name/SpaceX
//select * from company where name = 'SpaceX';
func searchByName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
	pathParams := mux.Vars(r)
	fmt.Fprintln(w, pathParams)
	if name, ok := pathParams["name"]; ok {
		row := db.QueryRow("SELECT * FROM companies WHERE name = $1", name)

		err := row.Scan(&name)
		switch {
		case err == sql.ErrNoRows:
			http.NotFound(w, r)
			return
		case err != nil:
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "name:", name)

	}

	//
	//
	//bk := Book{}
	//err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	//switch {
	//case err == sql.ErrNoRows:
	//	http.NotFound(w, r)
	//	return
	//case err != nil:
	//	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	//	return
	//}

}

//func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte(`{"message": "hello world"}`))
//}
