package routes

import (
	"net/http"
	"os"

	"github.com/danielsugianto/alterra-agmc-day3/controllers"
	m "github.com/danielsugianto/alterra-agmc-day3/middlewares"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New() *echo.Echo {
	keySecret := os.Getenv("JWT_SECRET")
	if keySecret == "" {
		panic("JWT Secret Missing")
	}
	e := echo.New()
	//call validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// call log middleware
	m.LogMiddlewares(e)
	v1 := e.Group("/v1")
	v1.POST("/login", controllers.LoginUsersController)

	// Unauthenticated URL
	// Unauthenticated Books
	v1.GET("/books", controllers.GetBooksController)
	v1.GET("/books/:id", controllers.GetBookController)

	// Unauthenticated Users
	v1.POST("/users", controllers.CreateUserController)

	//Authenticated URL
	authenticatedV1 := e.Group("/v1")
	//call middleware JWT
	authenticatedV1.Use(middleware.JWT([]byte(keySecret)))
	//authenticated users
	authenticatedV1.GET("/users", controllers.GetUsersController)
	authenticatedV1.GET("/users/:id", controllers.GetUserController)
	authenticatedV1.DELETE("/users/:id", controllers.DeleteUserController)
	authenticatedV1.PUT("/users/:id", controllers.UpdateUserController)
	//authenticated books
	authenticatedV1.POST("/books", controllers.CreateBookController)
	authenticatedV1.DELETE("/books/:id", controllers.DeleteBookController)
	authenticatedV1.PUT("/books/:id", controllers.UpdateBookController)
	return e
}
