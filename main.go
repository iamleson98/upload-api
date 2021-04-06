package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leminhson2398/zipper/modules/routes"
	"github.com/leminhson2398/zipper/modules/setting"
)

func main() {
	api := routes.API()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", setting.Port), api))
}
