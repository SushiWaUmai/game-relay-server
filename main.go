package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SushiWaUmai/game-server/env"
	"github.com/SushiWaUmai/game-server/server"
)

func main() {
	env.SetupDotenv()
	var router = server.SetupRoutes()

	log.Printf("Listening on Port %d...\n", env.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.PORT), router))
}
