package handlers

import "github.com/labstack/echo"

type UserHandler interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type UserHandlerImpl struct {}

func NewUserHandler() UserHandler {
	return &UserHandlerImpl{}
}

func (uh *UserHandlerImpl) 	Create(c echo.Context) error {
	return nil
}

func (uh *UserHandlerImpl) 	Read(c echo.Context) error {
	return nil
}

func (uh *UserHandlerImpl) 	Update(c echo.Context) error {
	return nil
}

func (uh *UserHandlerImpl) 	Delete(c echo.Context) error {
	return nil
}