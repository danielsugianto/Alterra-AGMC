package controllers

import (
	"net/http"
	"strconv"

	m "github.com/danielsugianto/alterra-agmc-day6/middlewares"
	"github.com/danielsugianto/alterra-agmc-day6/models"
	"github.com/labstack/echo/v4"
)

type usersController struct {
	usersUsecase models.UsersUsecase
}

// NewUsersUsecase will create new an usersController object representation of domain.UsersUsecase interface
func NewUsersController(usersUsecase models.UsersUsecase) models.UsersController {
	return &usersController{
		usersUsecase: usersUsecase,
	}
}

// get all users
func (usersController *usersController) GetUsersController(c echo.Context) error {
	users, err := usersController.usersUsecase.GetUsersUsecase()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// create new user
func (usersController *usersController) CreateUserController(c echo.Context) error {
	userParam := models.User{}
	c.Bind(&userParam)
	if err := c.Validate(userParam); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := usersController.usersUsecase.CreateUserUsecase(userParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// get user by ID
func (usersController *usersController) GetUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	_, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "Can't Get User")
	}
	user, err := usersController.usersUsecase.GetUserUsecase(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// delete user by ID
func (usersController *usersController) DeleteUserController(c echo.Context) error {
	id := c.Param("id")

	authId := m.ExtractTokenUserId(c)
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, "Can't Delete User")
	}
	if authId != intId {
		return c.JSON(http.StatusForbidden, "Can't Delete User")
	}
	errDelete := usersController.usersUsecase.DeleteUserUsecase(id)
	if errDelete != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by ID
func (usersController *usersController) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	authId := m.ExtractTokenUserId(c)
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, "Cant Update User")
	}
	if authId != intId {
		return c.JSON(http.StatusForbidden, "Cant Update User")
	}
	userParam := models.User{}
	c.Bind(&userParam)
	var user models.User

	user, errUpdate := usersController.usersUsecase.UpdateUserUsecase(id, userParam)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

// Login Controller
func (usersController *usersController) LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, e := usersController.usersUsecase.LoginUsersUsecase(user)
	if e != nil {
		return c.JSON(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}
