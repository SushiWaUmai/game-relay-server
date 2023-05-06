package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SushiWaUmai/relayroom/api"
	"github.com/SushiWaUmai/relayroom/env"
)

func main() {
	var router = api.SetupRoutes()

	log.Printf("Listening on Port %d...\n", env.PORT)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", env.PORT), router))
}
