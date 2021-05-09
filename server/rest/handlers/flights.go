package handlers

import (
	"net/http"

	"github.com/edualb/rest-graphql-grpc/server/rest/models"
	"github.com/labstack/echo"
)

type FlightsHandler interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	Update(c echo.Context) error
}

type FlightsHandlerImpl struct {
	db models.FlightsDB
}

func NewFlightsHandler() FlightsHandler {
	return &FlightsHandlerImpl{
		db: models.NewFlightsDB(),
	}
}

func (uh *FlightsHandlerImpl) Create(c echo.Context) error {
	var f models.Flights
	if err := c.Bind(&f); err != nil {
		return err
	}

	if err := uh.db.CreateFlight(f); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, f)
}

func (uh *FlightsHandlerImpl) Read(c echo.Context) error {
	id := c.Param("id")
	f, err := uh.db.FindFlightByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, f)
}

func (uh *FlightsHandlerImpl) Update(c echo.Context) error {
	var f models.Flights
	if err := c.Bind(&f); err != nil {
		return err
	}

	id := c.Param("id")

	if err := uh.db.UpdateFlightByID(id, f); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, f)
}
