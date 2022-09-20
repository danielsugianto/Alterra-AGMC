package routes

import (
	"github.com/danielsugianto/alterra-agmc-day6/config"
	"github.com/danielsugianto/alterra-agmc-day6/models"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

var (
	DB *gorm.DB
)

func init() {
	config.InitialDB()
}
func UnauthenticatedUserRoutes(e *echo.Group, usersController models.UsersController) {
	e.POST("/login", usersController.LoginUsersController)
	e.POST("/users", usersController.CreateUserController)
}

func AuthenticatedUserRoutes(e *echo.Group, usersController models.UsersController) {
	e.GET("/users", usersController.GetUsersController)
	e.GET("/users/:id", usersController.GetUserController)
	e.DELETE("/users/:id", usersController.DeleteUserController)
	e.PUT("/users/:id", usersController.UpdateUserController)
}
