package routes

import (
	"github.com/danielsugianto/alterra-agmc-day2/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")

	v1.GET("/books", controllers.GetBooksController)
	v1.POST("/books", controllers.CreateBookController)
	v1.GET("/books/:id", controllers.GetBookController)
	v1.DELETE("/books/:id", controllers.DeleteBookController)
	v1.PUT("/books/:id", controllers.UpdateBookController)

	v1.GET("/users", controllers.GetUsersController)
	v1.POST("/users", controllers.CreateUserController)
	v1.GET("/users/:id", controllers.GetUserController)
	v1.DELETE("/users/:id", controllers.DeleteUserController)
	v1.PUT("/users/:id", controllers.UpdateUserController)

	return e
}
