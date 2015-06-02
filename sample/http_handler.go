package http

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}
