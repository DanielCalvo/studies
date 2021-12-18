package main

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	infoLogger.Println("Hello world!")
	errorLogger.Println("Something went wrong!")
}

func someFunc() {
	infoLogger.Println("I am using infologger outside of main, this is why it is important for it to be a global variable!")
}
