package main

import (
	"github.com/edualb/rest-graphql-grpc/server/rest/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	uh := handlers.NewUserHandler()

	e.GET("/api/users", uh.Read)
}