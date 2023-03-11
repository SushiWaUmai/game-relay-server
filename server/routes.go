package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func heathcheck(c *gin.Context) {
	c.String(http.StatusOK, "Hello, API!")
}

type CreateOrJoinLobbyRequest struct {
	PlayerId string `json:"playerId"`
}

type CreateOrJoinLobbyResponse struct {
	LobbyId string `json:"lobbyId"`
}

func createOrJoinLobby(c *gin.Context) {
	var requestBody CreateOrJoinLobbyRequest

	if err := c.BindJSON(&requestBody); err != nil {
	}

	lobbyId := RandSeq(5)
	playerId := requestBody.PlayerId

	ip := c.Request.RemoteAddr

	err := RedisClient.SAdd("lobbies:"+lobbyId, playerId).Err()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = RedisClient.Set("player:"+playerId+":ip", ip, 0).Err()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	responseBody := CreateOrJoinLobbyResponse{
		LobbyId: lobbyId,
	}

	c.JSON(http.StatusOK, responseBody)
}

func getLobbies(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

func websocket(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("Trying to access lobby with id: %s...", id)

	c.String(http.StatusNotImplemented, "Not Implemented")
}

func SetupRoutes() *gin.Engine {
	setupRedis()

	router := gin.Default()
	router.GET("/", heathcheck)
	router.GET("/lobby", getLobbies)
	router.POST("/lobby", createOrJoinLobby)
	router.GET("/lobby/{id}", websocket)

	return router
}
