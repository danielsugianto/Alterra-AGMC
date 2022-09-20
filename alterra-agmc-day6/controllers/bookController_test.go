package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	actualBookIdString    = "1"
	notExistsBookIdString = "asd"
)

func TestCreateBookControllerValid(t *testing.T) {
	//body request
	bookJSON := `{
		"id":2,
		"name":"bebas",
		"year": 1998
	}`
	//setup
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateBookController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, d, rec.Body.String())
	}

}
func TestCreateBookControllerInvalid(t *testing.T) {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	//body request
	bookJSONInvalid := `{
		"id":"2",
		"name":"",
		"year": "1234"
	}`
	req := httptest.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(bookJSONInvalid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateBookController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetBooksControllerValid(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookControllerValid(t *testing.T) {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(actualBookIdString)

	// Assertions
	if assert.NoError(t, GetBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookControllerInvalid(t *testing.T) {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(notExistsBookIdString)

	// Assertions
	if assert.NoError(t, GetBookController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateBookControllerValid(t *testing.T) {

	e := echo.New()
	//body request
	updateBookJSON := `{
		"id":2,
		"name":"bebas",
		"year": 1998
	}`
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, "/v1/books", strings.NewReader(updateBookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(actualBookIdString)

	// Assertions
	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookControllerInvalid(t *testing.T) {

	e := echo.New()
	//body request
	updateBookJSON := `{
		"id":2,
		"name":"",
		"year": 1998
	}`
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, "/v1/books", strings.NewReader(updateBookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(notExistsBookIdString)

	// Assertions
	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestDeleteBookControllerValid(t *testing.T) {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodDelete, "/v1/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(actualBookIdString)
	// Assertions
	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBookControllerInvalid(t *testing.T) {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, "/v1/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(notExistsBookIdString)

	// Assertions
	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
