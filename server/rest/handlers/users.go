package handlers

import (
	"net/http"

	"github.com/edualb/rest-graphql-grpc/server/rest/models"
	"github.com/labstack/echo"
)

type UserHandler interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

var users *models.Users

type UserHandlerImpl struct{}

func NewUserHandler() UserHandler {
	return &UserHandlerImpl{}
}

func (uh *UserHandlerImpl) Create(c echo.Context) error {
	var u models.Users
	if err := c.Bind(&u); err != nil {
		return err
	}
	users = &u
	return c.JSON(http.StatusOK, users)
}

func (uh *UserHandlerImpl) Read(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func (uh *UserHandlerImpl) Update(c echo.Context) error {
	var u models.Users
	if err := c.Bind(&u); err != nil {
		return err
	}
	users = &u
	return c.JSON(http.StatusOK, users)
}

func (uh *UserHandlerImpl) Delete(c echo.Context) error {
	users = nil
	return c.JSON(http.StatusOK, users)
}
