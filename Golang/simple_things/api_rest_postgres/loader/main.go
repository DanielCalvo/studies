package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Company,Valuation ($B),Date Joined,Country,City,Industry,Select Inverstors,Founded Year,Total Raised,Financial Stage,Investors Count,Deal Terms,Portfolio Exits
type Company struct {
	Name       string
	Valuation  float64
	DateJoined string
	//	Country string
	//	City string
	//	Industry string
	//	SelectInvestors string
	//	TotalRaised int
	//	FinancialStage string
	//	InvestorCount int
	//	DealTerms int
	//	PortfolioExits int
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
	filename := "/home/daniel/Downloads/dataset/Unicorn_Companies.csv" //Tis a temporary location for development
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//var companies []Company

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		dt, err := time.Parse("1/2/2006", row[2])
		if err != nil {
			fmt.Println("Unable to time.Parse() provided date:", err)
			continue
		}

		company := Company{
			Name:       row[0],
			Valuation:  TrimValuation(row[1]),
			DateJoined: dt.Format("2006-01-02"),
		}

		fmt.Println(company)

		_, err = db.Exec("INSERT INTO company (NAME, VALUATION, DATEJOINED) VALUES ($1, $2, $3)", company.Name, company.Valuation, company.DateJoined)
		if err == nil {
			fmt.Println("Success inserting into the db:", company)
		}
		if err != nil {
			fmt.Println("Failed to insert into the db:", company)
		}
	}
}

func TrimValuation(s string) float64 {
	s = strings.Trim(s, "$")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1
	}
	return f
}
