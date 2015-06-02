package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// START OMIT
func BeerList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(beers)
}

func GetBeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, getBeer(vars["beer"]))

	//for _, beer := range beers {
	//	if beer.Id == vars["beer"] {
	//		json.NewEncoder(w).Encode(beer)
	//	}
	//}

}

// END OMIT
