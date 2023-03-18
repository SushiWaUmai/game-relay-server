package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func heathcheck(c *gin.Context) {
	c.String(http.StatusOK, "Hello, API!")
}

type joinLobbyRequest struct {
	LobbyId  string `json:"lobbyId"`
	PlayerId string `json:"playerId"`
}

type joinLobbyResponse struct {
}

type createLobbyResponse struct {
	LobbyId string `json:"lobbyId"`
}

func createLobby(c *gin.Context) {
	lobbyId := RandSeq(5)

	// Create Lobby in Redis
	res, err := RedisJsonHandler.JSONArrAppend("lobbies", ".", lobbyId)
	CheckRedisError(res, err)

	responseBody := createLobbyResponse{
		LobbyId: lobbyId,
	}

	c.JSON(http.StatusOK, responseBody)
}

func joinLobby(c *gin.Context) {
	var requestBody joinLobbyRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	playerId := requestBody.PlayerId
	lobbyId := requestBody.LobbyId

	ip := c.Request.RemoteAddr

	// Save Lobby in Redis
	res, err := RedisJsonHandler.JSONArrAppend(lobbyId, ".", playerId)
	CheckRedisError(res, err)

	// Save the player in Redis
	res, err = RedisJsonHandler.JSONArrAppend(playerId, ".", ip)
	CheckRedisError(res, err)

	responseBody := joinLobbyResponse{}

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
	router := gin.Default()
	router.GET("/", heathcheck)
	router.GET("/lobby", getLobbies)
	router.POST("/lobby/create", createLobby)
	router.POST("/lobby/join", joinLobby)
	router.GET("/lobby/{id}", websocket)

	return router
}
