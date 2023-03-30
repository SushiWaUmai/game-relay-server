package api

import (
	"log"
	"net/http"

	"github.com/SushiWaUmai/game-relay-server/game"
	"github.com/gin-gonic/gin"
)

func heathcheck(c *gin.Context) {
	c.String(http.StatusOK, "Hello, API!")
}

type createLobbyResponse struct {
	JoinCode string `json:"joinCode"`
}

func createLobby(c *gin.Context) {
	// Create Lobby
	lobby := game.NewLobby()

	responseBody := createLobbyResponse{
		JoinCode: lobby.JoinCode,
	}

	c.JSON(http.StatusOK, responseBody)
}

func getLobbies(c *gin.Context) {
	lobbies := make([]*game.Lobby, 0)

	game.Lobbies.Range(func(key, value interface {}) bool {
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
		log.Fatal("Could not find lobby")
		c.String(http.StatusNotFound, "Could not find lobby")
		return
	}

	lobby := value.(*game.Lobby)

	lobby.ServeHTTP(c.Writer, c.Request)
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/", heathcheck)
	router.GET("/lobby", getLobbies)
	router.POST("/lobby", createLobby)
	router.GET("/lobby/:joinCode", joinLobby)

	return router
}
