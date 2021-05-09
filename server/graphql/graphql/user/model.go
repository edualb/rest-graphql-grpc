package user

import gql "github.com/graphql-go/graphql"

type User struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	FlightIDs []string `json:"flight_ids,omitempty"`
}

var userType = gql.NewObject(
	gql.ObjectConfig{
		Name: "User",
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
