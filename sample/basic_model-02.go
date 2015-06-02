package main

type Beer struct {
	Name        string  `json:"name"`
	Id          string  `json:"id"`
	Img         string  `json:"img"`
	Description string  `json:"description"`
	Alcohol     float32 `json:"alcohol"`
}

type Beers []Beer
