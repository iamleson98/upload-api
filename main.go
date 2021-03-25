package main

import (
	"log"
	"net/http"

	"github.com/leminhson2398/zipper/modules/routes"
)

func main() {
	api := routes.API()

	log.Fatal(http.ListenAndServe(":8000", api))
}
