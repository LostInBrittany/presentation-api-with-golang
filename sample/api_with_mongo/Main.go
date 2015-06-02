package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	openDbConnection()

	loadBeers()

	log.Fatal(http.ListenAndServe(":7777", router))
}
