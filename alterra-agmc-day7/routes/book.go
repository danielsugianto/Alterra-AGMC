package routes

import (
	"github.com/danielsugianto/alterra-agmc-day7/controllers"
	"github.com/labstack/echo/v4"
)

func UnauthenticatedBookRoutes(e *echo.Group) {
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
}

func AuthenticatedBookRoutes(e *echo.Group) {
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
}
