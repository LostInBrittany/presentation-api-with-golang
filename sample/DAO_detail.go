
var session mgo.Session
var err error

func openDbConnection() {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
    panic(err)
  }
  defer session.Close()
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
