package main

import (
	"log"
	"net/http"
	api "tripmemories/server"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":3333"

	router := mux.NewRouter()

	api.SetupRoutes(router)

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
