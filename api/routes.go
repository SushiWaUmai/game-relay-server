package api

import (
	"fmt"
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
	var lobbies []game.Lobby

	game.Lobbies.Range(func(key any, value any) bool {
		l := value.(game.Lobby)
		_ = append(lobbies, l)

		return true
	})

	c.JSON(http.StatusOK, lobbies)
}

func joinLobby(c *gin.Context) {
	joinCode := c.Param("joinCode")
	fmt.Printf("Trying to access lobby with joinCode: %s...", joinCode)

	value, ok := game.Lobbies.Load(joinCode)
	if !ok {
		c.String(http.StatusInternalServerError, "Could not find lobby")
	}
	lobby := value.(game.Lobby)

	lobby.ServeHTTP(c.Writer, c.Request)
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/", heathcheck)
	router.GET("/lobby", getLobbies)
	router.POST("/lobby", createLobby)
	router.GET("/lobby/{joinCode}", joinLobby)

	return router
}
