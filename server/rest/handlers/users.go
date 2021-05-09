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

type UserHandlerImpl struct {
	db models.UsersDB
}

func NewUserHandler() UserHandler {
	return &UserHandlerImpl{
		db: models.NewUsersDB(),
	}
}

func (uh *UserHandlerImpl) Create(c echo.Context) error {
	var u models.Users
	if err := c.Bind(&u); err != nil {
		return err
	}

	if err := uh.db.CreateUser(u); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, u)
}

func (uh *UserHandlerImpl) Read(c echo.Context) error {
	id := c.Param("id")
	u, err := uh.db.FindUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func (uh *UserHandlerImpl) Update(c echo.Context) error {
	var u models.Users
	if err := c.Bind(&u); err != nil {
		return err
	}

	id := c.Param("id")

	if err := uh.db.UpdateUser(id, u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}

func (uh *UserHandlerImpl) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uh.db.DeleteUserByID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, id)
}
