package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// START OMIT
func BeerList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(beers)
}

func GetBeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for _, beer := range beers {
		if beer.Id == vars["beer"] {
			json.NewEncoder(w).Encode(beer)
		}
	}

}

// END OMIT
