package main

type Beer struct {
	Name        string
	Id          string
	Img         string
	Description string
	Alcohol     float32
}

type Beers []Beer

// START OMIT

beers := Beers {
  Beer { 
    Alcohol: 6.8
    Description: "The classic blonde abbey ale, with a gentle roundness, low on bitterness.",
    Id: "AffligemBlond",
    Img: "img/AffligemBlond.jpg",
    Name: "Affligem Blond" 
  },
  
  Beer {
    Alcohol: 9.2
    Description: "A dry but rich flavoured beer with complex fruity and spicy flavours.",
    Id: "TrappistesRochefort8",
    Img: "img/TrappistesRochefort8.jpg",
    Name: "Rocherfort 8"     
  }
}

// END OMIT