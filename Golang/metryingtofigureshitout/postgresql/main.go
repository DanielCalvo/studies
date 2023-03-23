package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

/*
CREATE TABLE MEMES (
   NAME TEXT NOT NULL,
   FUNNY BOOL NOT NULL
);

INSERT INTO memes (name, funny) VALUES ('pepe', TRUE);
*/

type Meme struct {
	Name  string
	Funny bool
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:example@localhost?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connection to database established!")
}

// Todo here: Insert some records in a table, select them back with a select+where statement
// Bonus: Marshall the results to json and print them to stdout
func main() {
	name := "pepe"
	funny := true

	_, err := db.Exec("INSERT INTO memes (name, funny) VALUES ($1, $2)", name, funny)
	if err != nil {
		log.Fatalln("Unable to insert into memes table:", err)
	}
	log.Println("Meme inserted on table!")

	//Now lets retrieve some meme information and print them as json!
	rows, err := db.Query("SELECT * FROM memes WHERE name = $1", "pepe")
	if err != nil {
		log.Fatalln("Unable to run select query:", err)
	}

	var memes []Meme

	for rows.Next() {
		var m Meme
		err = rows.Scan(&m.Name, &m.Funny)
		switch {
		case err == sql.ErrNoRows:
			log.Println("No rows match query :(", err)
			continue
		case err != nil:
			log.Println("SQL error:", err)
			continue
		}
		log.Println(m.Name, m.Funny)
		memes = append(memes, m)
	}
	fmt.Println(memes)

	b, err := json.Marshal(memes)
	if err != nil {
		log.Fatalln("Unable to mashall memes into json:", err)
	}
	fmt.Println(string(b))

}
