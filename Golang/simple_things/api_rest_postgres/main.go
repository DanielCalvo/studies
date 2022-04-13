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

type Company struct {
	Id         int
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
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})

	api.HandleFunc("/companies/name/{name}", searchByName).Methods(http.MethodGet)
	api.HandleFunc("/company", addCompany).Methods(http.MethodPost)
	api.HandleFunc("/companies/name/{name}", deleteByName).Methods(http.MethodDelete)
	api.HandleFunc("/companies/name/{name}", updateByName).Methods(http.MethodPut)
	log.Fatalln(http.ListenAndServe(":8080", r))

}

/*
Case sensitive:
curl -X GET http://localhost:8080/api/v1/companies/name/SpaceX
*/
func searchByName(w http.ResponseWriter, r *http.Request) {
	var c Company

	pathParams := mux.Vars(r)
	if name, ok := pathParams["name"]; ok {
		row := db.QueryRow("SELECT * FROM company WHERE name = $1", name)

		err := row.Scan(&c.Id, &c.Name, &c.Valuation, &c.DateJoined)
		switch {
		case err == sql.ErrNoRows:
			log.Println("Company not found in db:", pathParams["name"])
			http.NotFound(w, r)
			return
		case err != nil:
			log.Println("Error querying DB:", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(c)
		if err != nil {
			log.Println("Error marshalling json:", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
		fmt.Fprintln(w, string(b))
		log.Println("Search successful for:", c.Name)
	}
}

/*
curl -X POST http://localhost:8080/api/v1/company -H 'Content-Type: application/json' -d '{"Name":"memecorp","Valuation":1.1, "DateJoined":"2002/12/01"}'
*/
func addCompany(w http.ResponseWriter, r *http.Request) {
	var c Company
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO company (name, valuation, datejoined) VALUES ($1, $2, $3)", c.Name, c.Valuation, c.DateJoined)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Println("Error inserting into db:", err)
		return
	}
	log.Println("Added to database:", c)

}

/*
curl -X DELETE http://localhost:8080/api/v1/companies/name/memecorp
*/
func deleteByName(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	name, ok := queries["name"]
	if ok {
		_, err := db.Exec("DELETE FROM company WHERE name=$1;", name)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		log.Println("Deleted company:", name)
	}
}

/*
Generally, in practice, use PUT for UPDATE operations. Always use POST for CREATE operations
curl -X PUT http://localhost:8080/api/v1/companies/name/memecorp -H 'Content-Type: application/json' -d '{"Name":"memecorp","Valuation":2.2, "DateJoined":"2002/12/01"}'
*/
func updateByName(w http.ResponseWriter, r *http.Request) {
	var c Company
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("update called!")
	queries := mux.Vars(r)
	name, ok := queries["name"]
	fmt.Println(name, ok)
	if ok {
		_, err := db.Exec("UPDATE company SET name=$1, valuation=$2, datejoined=$3 WHERE name=$1;", c.Name, c.Valuation, c.DateJoined)
		if err != nil {
			log.Println("Error updating company:", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	}
	log.Println("Updated:", name, "with the following:", c)
}
