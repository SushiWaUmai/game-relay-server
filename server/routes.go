package server

import (
	"fmt"
	"net/http"

	"github.com/SushiWaUmai/game-relay-server/db"
	"github.com/gin-gonic/gin"
)

func heathcheck(c *gin.Context) {
	c.String(http.StatusOK, "Hello, API!")
}

type createLobbyResponse struct {
	JoinCode string `json:"joinCode"`
}

func createLobby(c *gin.Context) {
	joinCode := RandSeq(5)

	// Create Lobby
	db.DatabaseConnection.Create(&db.Lobby{
		JoinCode: joinCode,
	})

	responseBody := createLobbyResponse{
		JoinCode: joinCode,
	}

	c.JSON(http.StatusOK, responseBody)
}

func getLobbies(c *gin.Context) {
	var lobbies []db.Lobby
	db.DatabaseConnection.Find(&lobbies)

	c.JSON(http.StatusOK, lobbies)
}

func joinLobby(c *gin.Context) {
	joinCode := c.Param("joinCode")
	fmt.Printf("Trying to access lobby with joinCode: %s...", joinCode)

	ip := c.Request.RemoteAddr

	var lobby db.Lobby
	db.DatabaseConnection.Where("JoinCode = ?", joinCode).First(&lobby)

	// Save the player
	db.DatabaseConnection.Create(&db.Player{
		LobbyID: lobby.ID,
		IP:      ip,
	})

	c.String(http.StatusNotImplemented, "Not Implemented")
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/", heathcheck)
	router.GET("/lobby", getLobbies)
	router.POST("/lobby", createLobby)
	router.GET("/lobby/{joinCode}", joinLobby)

	return router
}
