package controllers

import (
	"net/http"
	"strconv"

	"github.com/danielsugianto/alterra-agmc-day4/lib/database"
	m "github.com/danielsugianto/alterra-agmc-day4/middlewares"
	"github.com/danielsugianto/alterra-agmc-day4/models"
	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	users, err := database.GetUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := database.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// get user by ID
func GetUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	_, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "Can't Get User")
	}
	user, err := database.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// delete user by ID
func DeleteUserController(c echo.Context) error {
	id := c.Param("id")

	authId := m.ExtractTokenUserId(c)
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, "Can't Delete User")
	}
	if authId != intId {
		return c.JSON(http.StatusForbidden, "Can't Delete User")
	}
	errDelete := database.DeleteUser(id)
	if errDelete != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by ID
func UpdateUserController(c echo.Context) error {
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

	user, errUpdate := database.UpdateUser(id, userParam)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

// Login Controller
func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, e := database.LoginUser(&user)
	if e != nil {
		return c.JSON(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}
