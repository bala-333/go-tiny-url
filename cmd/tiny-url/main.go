package main

import (
	"go-tiny-url/endpoints"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.POST("/", endpoints.ShortUrl)
	log.Println("Listing for requests at http://localhost:8081/")
	log.Fatal(http.ListenAndServe(":8081", mux))
}