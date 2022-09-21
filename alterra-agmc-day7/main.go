package main

import (
	"github.com/danielsugianto/alterra-agmc-day7/config"
	"github.com/danielsugianto/alterra-agmc-day7/controllers"
	"github.com/danielsugianto/alterra-agmc-day7/lib/database"
	"github.com/danielsugianto/alterra-agmc-day7/routes"
	"github.com/danielsugianto/alterra-agmc-day7/usecase"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	MongoDB *mongo.Database
)

func main() {
	godotenv.Load(".env")

	DB = config.InitialDB()
	MongoDB = config.InitialMongoDB()
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := controllers.NewUsersController(usersUsecase)

	booksMySQLRepo := database.NewMongoDBBooksRepository(MongoDB)
	booksUsecase := usecase.NewBooksUsecase(booksMySQLRepo)
	booksController := controllers.NewBooksController(booksUsecase)
	e := routes.New(usersController, booksController)
	// start the server, and log if it fails
	e.Start(":8000")
}
