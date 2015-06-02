package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// START OMIT
var session mgo.Session
var err error

func openDbConnection() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
}

func createBeer(beer Beer) {
	c := session.DB("test").C("beers")
	err = c.Insert(beer)
	if err != nil {
		log.Fatal(err)
	}
}

func getBeer(id string) {
	c := session.DB("test").C("beers")
	result := Beer{}
	err = c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	jsonBeer, _ := json.Marshal(result)
	fmt.Println(string(jsonBeer))
}

// END OMIT
