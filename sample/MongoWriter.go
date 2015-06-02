package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Beer struct {
	Name        string  `json:"name"`
	Id          string  `json:"id"`
	Img         string  `json:"img"`
	Description string  `json:"description"`
	Alcohol     float32 `json:"alcohol"`
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
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("beers")
	err = c.Insert(beers[0], beers[1])
	if err != nil {
		log.Fatal(err)
	}

	result := Beer{}
	err = c.Find(bson.M{"name": "Rocherfort 8"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	jsonBeer, _ := json.Marshal(beers[1])

	fmt.Println("Beer:", string(jsonBeer))
}

// END OMIT
