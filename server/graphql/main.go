package main

import (
	"github.com/edualb/rest-graphql-grpc/server/graphql/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	h := handler.New(&handler.Config{
		Schema:   graphql.NewSchema(),
		Pretty:   true,
		GraphiQL: true,
	})

	e.GET("/graphql", echo.WrapHandler(h))
	e.POST("/graphql", echo.WrapHandler(h))

	e.Logger.Fatal(e.Start(":1234"))
}
