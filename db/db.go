package db

import (
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DatabaseConnection *gorm.DB

type Lobby struct {
	gorm.Model
}

func init() {
	setupDatabase();
}

func setupDatabase() {
	var err error
	DatabaseConnection, err = gorm.Open(sqlite.Open("game.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to create database connection")
		os.Exit(1)
	}
}
