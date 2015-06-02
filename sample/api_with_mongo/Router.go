package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// START OMIT
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{"BeerList", "GET", "/beers", BeerList},
	Route{"GetBeer", "GET", "/beer/{beer}", GetBeer},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).Path(route.Pattern).
			Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

// END OMIT
