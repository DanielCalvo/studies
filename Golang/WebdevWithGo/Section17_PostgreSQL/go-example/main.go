package main

/*
psql -h localhost -U postgres
CREATE DATABASE bookstore;
CREATE USER dummy with password 'banana';
GRANT ALL PRIVILEGES ON database bookstore TO dummy;
ALTER USER dummy WITH superuser;

CREATE TABLE books (
  isbn    char(14)     PRIMARY KEY NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);

\d books

INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);

*/

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://dummy:banana@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, bk := range bks {
		// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}

}
