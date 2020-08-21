package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

var db *sql.DB
var err error

func main(){
	//db, err = sql.Open("mysql", "user:password@tcp(127.0.0.7:3306)/test02?charset=utf8")
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.7:3306)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query(`SELECT name FROM person;`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var name string
	for rows.Next(){
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
	fmt.Println(name)
	}


	fmt.Println("done yo")
}
