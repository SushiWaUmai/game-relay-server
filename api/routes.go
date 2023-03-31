package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SushiWaUmai/game-relay-server/game"
	"github.com/gin-gonic/gin"
)

func heathcheck(c *gin.Context) {
	c.String(http.StatusOK, "Hello, API!")
}

func createLobby(c *gin.Context) {
	// Create Lobby
	lobby := game.NewLobby()

	responseBody := game.Lobby{
		JoinCode: lobby.JoinCode,
	}

	c.JSON(http.StatusOK, responseBody)
}

func getLobbies(c *gin.Context) {
	lobbies := make([]*game.Lobby, 0)

	game.Lobbies.Range(func(key, value interface{}) bool {
		l := value.(*game.Lobby)
		lobbies = append(lobbies, l)

		return true
	})

	c.JSON(http.StatusOK, lobbies)
}

func joinLobby(c *gin.Context) {
	joinCode := c.Param("joinCode")
	log.Printf("Trying to access lobby with joinCode: %s...\n", joinCode)

	value, ok := game.Lobbies.Load(joinCode)

	if !ok {
		log.Println("Could not find lobby")
		c.String(http.StatusNotFound, "Could not find lobby")
		return
	}

	lobby := value.(*game.Lobby)

	lobby.ServeHTTP(c.Writer, c.Request)
}

func getClients(c *gin.Context) {
	joinCode := c.Param("joinCode")

	value, ok := game.Lobbies.Load(joinCode)

	if !ok {
		log.Println("Could not find lobby")
		c.String(http.StatusNotFound, "Could not find lobby")
		return
	}

	lobby := value.(*game.Lobby)
	clients := lobby.Clients()
	c.JSON(http.StatusOK, clients)
}

func getClient(c *gin.Context) {
	joinCode := c.Param("joinCode")

	value, ok := game.Lobbies.Load(joinCode)

	if !ok {
		log.Println("Could not find lobby")
		c.String(http.StatusNotFound, "Could not find lobby")
		return
	}
	lobby := value.(*game.Lobby)

	clientId, err := strconv.Atoi(c.Param("clientId"))

	if err != nil {
		log.Println("Failed to parse clientId")
		c.String(http.StatusNotFound, "Failed to parse clientId")
		return
	}

	client := lobby.GetClient(uint(clientId))
	c.JSON(http.StatusOK, client)
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/", heathcheck)
	router.GET("/lobby", getLobbies)
	router.POST("/lobby", createLobby)
	router.GET("/lobby/:joinCode", joinLobby)
	router.GET("/lobby/:joinCode/clients", getClients)
	router.GET("/lobby/:joinCode/clients/:clientId", getClient)

	return router
}
