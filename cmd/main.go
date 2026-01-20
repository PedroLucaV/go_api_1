package main

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

var middlewares = []Middleware{
	TokenAuthMiddleware,
}

func main() {
	mux := http.NewServeMux()
	var handler http.HandlerFunc = handleClientProfile
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	mux.Handle("GET /user/profile/{clientId}", handler)

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
