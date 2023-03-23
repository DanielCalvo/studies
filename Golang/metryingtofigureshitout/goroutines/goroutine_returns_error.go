package main

import (
	"errors"
	"log"
	"time"
)

// I want to have a function that runs forever in a go routine and might eventually return an error
// If this go routine returns an error, I would like to stop the program
func main() {
	log.Println("let's go!")

	chErr := make(chan error)
	go functionThatRunsForever(chErr)

	//You can't handle the error channel in the main go routine, otherwise you'll block the program!
	go func(err chan error) {
		log.Fatalln("Got an error on the goroutine that runs forever:", <-err)
	}(chErr)

	//In here goes the webserver (think Prometheus exporter, so you could argue that this "time.Sleep" is a place holder for something line http.ListenAndServe()
	time.Sleep(time.Second * 100)
}

func functionThatRunsForever(chErr chan error) chan error {
	counter := 0
	for {
		log.Println("I'm here happily running forever!")
		time.Sleep(time.Second)
		counter++
		if counter == 10 {
			chErr <- errors.New("Something went mega wrong")
			return chErr
		}
	}
}
