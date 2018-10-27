package main

import (
	"github.com/eexe1/supermaid-go/network/routes"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {

	router := routes.Routes()

	walkFunc := func(method string, route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler) error {
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8080", router))

}
