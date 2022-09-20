package main

import (
	"github.com/danielsugianto/alterra-agmc-day7/config"
	"github.com/danielsugianto/alterra-agmc-day7/controllers"
	"github.com/danielsugianto/alterra-agmc-day7/lib/database"
	"github.com/danielsugianto/alterra-agmc-day7/routes"
	"github.com/danielsugianto/alterra-agmc-day7/usecase"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	DB = config.InitialDB()
}

func main() {
	godotenv.Load(".env")

	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := controllers.NewUsersController(usersUsecase)
	e := routes.New(usersController)
	// start the server, and log if it fails
	e.Start(":8000")
}
