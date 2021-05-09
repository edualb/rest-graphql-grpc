package user

import (
	gql "github.com/graphql-go/graphql"
)

var usersType = gql.NewObject(
	gql.ObjectConfig{
		Name: "Users",
		Fields: gql.Fields{
			"id": &gql.Field{
				Type: gql.String,
			},
			"name": &gql.Field{
				Type: gql.String,
			},
			"flight_ids": &gql.Field{
				Type: gql.NewList(gql.String),
			},
		},
	},
)
