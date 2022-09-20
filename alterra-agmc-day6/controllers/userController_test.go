package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"testing"

	"github.com/danielsugianto/alterra-agmc-day6/config"
	"github.com/danielsugianto/alterra-agmc-day6/lib/database"
	m "github.com/danielsugianto/alterra-agmc-day6/middlewares"
	"github.com/danielsugianto/alterra-agmc-day6/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	actualUserIdString    = "1"
	actualUserIdInt, _    = strconv.Atoi(actualUserIdString)
	notExistsUserIdString = "asd"
	tokenActualUser, _    = m.CreateToken(actualUserIdInt)
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

var (
	DB *gorm.DB
)

func TestCreateUserControllerValid(t *testing.T) {
	//init DB
	DB = config.InitialDB()

	//body request
	userJSON := `{
		"name":"test test test",
		"email":"popoaa@testtest.com",
		"password": "1283948149"
	}`
	//setup
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, usersController.CreateUserController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, d, rec.Body.String())
	}

}
func TestCreateUserControllerInvalid(t *testing.T) {
	DB = config.InitialDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	userJSONInvalid := `{
		"name":"",
		"email":"test@mailinator.com",
		"password": "1283948149"
	}`
	req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(userJSONInvalid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, usersController.CreateUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
func TestLoginControllerValid(t *testing.T) {
	DB = config.InitialDB()
	userJSON := `{		
		"email":"test@mailinator.com",
		"password": "1234"
	}`

	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, usersController.LoginUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLoginControllerInvalid(t *testing.T) {
	DB = config.InitialDB()
	userJSON := `{		
		"email":"test@",
		"password": "1234"
	}`
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, usersController.LoginUsersController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetUsersControllerValid(t *testing.T) {
	DB = config.InitialDB()
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, usersController.GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserControllerValid(t *testing.T) {
	DB = config.InitialDB()
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(actualUserIdString)

	// Assertions
	if assert.NoError(t, usersController.GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserControllerInvalid(t *testing.T) {
	DB = config.InitialDB()
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(notExistsUserIdString)

	// Assertions
	if assert.NoError(t, usersController.GetUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateUserControllerValid(t *testing.T) {
	DB = config.InitialDB()
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	updateUserJSON := `{
		"name":"test test test",
		"email":"popoaa@testtest.com",
		"password": "1283948149"
	}`
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, "/v1/users/", strings.NewReader(updateUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(actualUserIdString)
	c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenActualUser))

	// Assertions
	if assert.NoError(t, usersController.UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateUserControllerInvalid(t *testing.T) {
	DB = config.InitialDB()
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	updateUserJSON := `{
		"name":"test test test",
		"email":"popoaa@testtest.com",
		"password": "1283948149"
	}`
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, "/v1/users/", strings.NewReader(updateUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(notExistsUserIdString)
	c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenActualUser))

	// Assertions
	if assert.NoError(t, usersController.UpdateUserController(c)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}

func TestDeleteUserControllerValid(t *testing.T) {
	DB = config.InitialDB()
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodDelete, "/v1/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(actualUserIdString)
	c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenActualUser))

	// Assertions
	if assert.NoError(t, usersController.DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUserControllerInvalid(t *testing.T) {
	DB = config.InitialDB()
	usersMySQLRepo := database.NewMySQLUsersRepository(DB)
	usersUsecase := usecase.NewUsersUsecase(usersMySQLRepo)
	usersController := NewUsersController(usersUsecase)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, "/v1/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(notExistsUserIdString)
	c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenActualUser))

	// Assertions
	if assert.NoError(t, usersController.DeleteUserController(c)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}
