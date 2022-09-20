package models

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token"`
}

type UsersUsecase interface {
	GetUsersUsecase() ([]User, error)
	CreateUserUsecase(userParam User) (User, error)
	GetUserUsecase(id string) (User, error)
	DeleteUserUsecase(id string) error
	UpdateUserUsecase(id string, userParam User) (User, error)
	LoginUsersUsecase(user User) (interface{}, error)
}

type UsersMySQLRepository interface {
	GetUsers() ([]User, error)
	CreateUser(userParam User) (User, error)
	GetUser(id string) (User, error)
	DeleteUser(id string) error
	UpdateUser(id string, userParam User) (User, error)
	LoginUser(user *User) (interface{}, error)
}

type UsersController interface {
	GetUsersController(c echo.Context) error
	CreateUserController(c echo.Context) error
	GetUserController(c echo.Context) error
	DeleteUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
	LoginUsersController(c echo.Context) error
}
