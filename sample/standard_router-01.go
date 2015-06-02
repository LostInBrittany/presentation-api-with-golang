http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
  log.Println("serving", req.URL)
  fmt.Fprintln(w, "Hello world")
})

