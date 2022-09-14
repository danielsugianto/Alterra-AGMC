package main

import (
	"github.com/danielsugianto/alterra-agmc-day3/config"
	"github.com/danielsugianto/alterra-agmc-day3/routes"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	config.InitialDB()
}

func main() {
	godotenv.Load(".env")

	e := routes.New()
	// start the server, and log if it fails
	e.Start(":8000")
}
