package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func heathcheck(c *gin.Context) {
	c.String(http.StatusOK, "Hello, API!")
}

func createLobby(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
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
	router.POST("/lobby", createLobby)
	router.GET("/lobby/{id}", websocket)

	return router
}
