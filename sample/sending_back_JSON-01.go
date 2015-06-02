package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Beer struct {
	Name        string
	Id          string
	Img         string
	Description string
	Alcohol     float32
}

type Beers []Beer

var beers = Beers{
	Beer{
		Alcohol:     6.8,
		Description: "Classic blonde abbey ale, with a gentle roundness, low on bitterness.",
		Id:          "AffligemBlond",
		Img:         "img/AffligemBlond.jpg",
		Name:        "Affligem Blond",
	},
	// [...]
	Beer{
		Alcohol:     9.2,
		Description: "A dry but rich flavoured beer with complex fruity and spicy flavours.",
		Id:          "TrappistesRochefort8",
		Img:         "img/TrappistesRochefort8.jpg",
		Name:        "Rocherfort 8",
	},
}

// START OMIT
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Println("Running on http://127.0.0.1:7777")
	log.Fatal(http.ListenAndServe(":7777", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(beers)
}

// END OMIT
