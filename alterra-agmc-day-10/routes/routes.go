package routes

import (
	"net/http"
	"os"

	m "github.com/danielsugianto/alterra-agmc-day-10/middlewares"
	"github.com/danielsugianto/alterra-agmc-day-10/models"
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

func New(usersController models.UsersController, booksController models.BooksController) *echo.Echo {
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
	UnauthenticatedUserRoutes(v1, usersController)
	UnauthenticatedBookRoutes(v1, booksController)

	//Authenticated URL
	authenticatedV1 := e.Group("/v1")
	//call middleware JWT
	authenticatedV1.Use(middleware.JWT([]byte(keySecret)))
	AuthenticatedUserRoutes(authenticatedV1, usersController)
	AuthenticatedBookRoutes(authenticatedV1, booksController)
	return e
}
