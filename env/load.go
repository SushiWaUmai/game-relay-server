package env

import (
	"os"
	"strconv"
)

var (
	PORT int
)

func loadEnv() {
	PORT, _ = strconv.Atoi(os.Getenv("PORT"))
}
