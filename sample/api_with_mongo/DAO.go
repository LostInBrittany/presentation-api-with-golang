package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"os"
)

var (
	mgoSession   *mgo.Session
	databaseName = "myDB"
	err          error
)

func openDbConnection() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err) // no, not really
		}
	}
	return mgoSession.Clone()
}

func loadBeers() {
	file, e := ioutil.ReadFile("./beers.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file))

	var beerList Beers
	json.Unmarshal(file, &beerList)
	fmt.Printf("Results: %v\n", beerList)

	for _, beer := range beerList {
		createBeer(beer)
	}
}

func createBeer(beer Beer) {
	c := mgoSession.DB("test").C("beers")
	err = c.Insert(beer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Beer %s created\n", beer.Name)
}

func getBeer(id string) string {
	c := mgoSession.DB("test").C("beers")
	result := Beer{}
	err = c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Println(err)
		return ""
	}
	jsonBeer, _ := json.Marshal(result)
	fmt.Println(string(jsonBeer))
	return string(jsonBeer)
}
