package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	http.Handle("/hello", String("Hello world"))
	fmt.Println("serving on http://localhost:7777/hello")
	http.ListenAndServe("localhost:7777", nil)
}
