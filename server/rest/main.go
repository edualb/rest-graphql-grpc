package main

import (
	"github.com/edualb/rest-graphql-grpc/server/rest/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	uh := handlers.NewUserHandler()

	e.GET("/api/users/:id", uh.Read)
	e.PUT("/api/users/:id", uh.Update)
	e.POST("/api/users", uh.Create)
	e.DELETE("/api/users/:id", uh.Delete)

	fh := handlers.NewFlightsHandler()

	e.GET("/api/flights/:id", fh.Read)
	e.PUT("/api/flights/:id", fh.Update)
	e.POST("/api/flights", fh.Create)

	e.Logger.Fatal(e.Start(":8080"))
}
