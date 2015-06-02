package main

// START OMIT
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
	//  [...]
	// END OMIT
	Beer{
		Alcohol:     9.2,
		Description: "A dry but rich flavoured beer with complex fruity and spicy flavours.",
		Id:          "TrappistesRochefort8",
		Img:         "img/TrappistesRochefort8.jpg",
		Name:        "Rocherfort 8",
	},
}
