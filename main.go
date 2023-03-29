package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SushiWaUmai/game-relay-server/env"
	"github.com/SushiWaUmai/game-relay-server/api"
)

func main() {
	var router = api.SetupRoutes()

	log.Printf("Listening on Port %d...\n", env.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.PORT), router))
}
