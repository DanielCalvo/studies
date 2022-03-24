package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

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

//Company,Valuation ($B),Date Joined,Country,City,Industry,Select Inverstors,Founded Year,Total Raised,Financial Stage,Investors Count,Deal Terms,Portfolio Exits
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

		company := Company{
			Name:      row[0],
			Valuation: TrimValuation(row[1]),
		}

		//Need to convert that B thing to numbers! Make a function for that!
		fmt.Println(company)

	}
}

/*
About the date: PostgreSQL uses the yyyy-mm-dd format for storing and inserting date values.
And golang can also do yyyy-mm-dd https://golang.cafe/blog/golang-time-format-example.html
Let's convert the data from the csv to yyyy-mm-dd format then!
*/

func TrimValuation(s string) float64 {
	s = strings.Trim(s, "$")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1
	}
	return f
}
