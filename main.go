package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SushiWaUmai/game-server/env"
	"github.com/SushiWaUmai/game-server/server"
	"github.com/SushiWaUmai/game-server/db"
)

func main() {
	var router = server.SetupRoutes()

	data := &db.Lobby{}
	log.Println(data)

	log.Printf("Listening on Port %d...\n", env.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.PORT), router))
}
