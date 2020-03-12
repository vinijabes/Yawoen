package main

import (
	"log"
	"net/http"

	"app/src/routes"
)

func main() {
	router := routes.NewRouter()
	log.Print("Iniciou")
	log.Fatal(http.ListenAndServe(":5000", router))
}
