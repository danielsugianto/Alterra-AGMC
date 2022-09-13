package main

import (
	"github.com/danielsugianto/alterra-agmc-day2/config"
	"github.com/danielsugianto/alterra-agmc-day2/routes"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	config.InitialDB()
}

func main() {
	e := routes.New()

	// start the server, and log if it fails
	e.Start(":8000")
}
