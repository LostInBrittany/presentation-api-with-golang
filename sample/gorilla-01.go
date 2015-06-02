package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// START OMIT
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/hello/{who}", HelloWho)
	log.Println("Running on http://127.0.0.1:7777")
	log.Fatal(http.ListenAndServe(":7777", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func HelloWho(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %s", vars["who"])
}

// END OMIT
