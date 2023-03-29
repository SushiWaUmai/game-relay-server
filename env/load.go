package env

import (
	"os"
	"strconv"
)

var (
	PORT                = 8080
	SOCKET_BUFFER_SIZE  = 1024
	MESSAGE_BUFFER_SIZE = 256
)

func loadEnv() {
	var err error
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		PORT = 8080
	}

	SOCKET_BUFFER_SIZE, err = strconv.Atoi(os.Getenv("SOCKET_BUFFER_SIZE"))
	if err != nil {
		SOCKET_BUFFER_SIZE = 1024
	}

	MESSAGE_BUFFER_SIZE, err = strconv.Atoi(os.Getenv("MESSAGE_BUFFER_SIZE"))
	if err != nil {
		MESSAGE_BUFFER_SIZE = 256
	}
}
