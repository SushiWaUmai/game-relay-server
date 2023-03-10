package main

import (
	"fmt"
	"log"
	"net/http"

	e "github.com/SushiWaUmai/game-server/env"
	r "github.com/SushiWaUmai/game-server/server"
)

func main() {
	e.SetupDotenv()
	var router = r.SetupRoutes()

	log.Printf("Listening on Port %d...\n", e.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", e.PORT), router))
}
