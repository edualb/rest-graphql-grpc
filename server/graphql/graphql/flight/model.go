package flight

import gql "github.com/graphql-go/graphql"

type Flight struct {
	ID           string   `json:"id"`
	From         string   `json:"from"`
	To           string   `json:"to"`
	PassengerIDs []string `json:"passenger_ids,omitempty"`
}

var userType = gql.NewObject(
	gql.ObjectConfig{
		Name: "Flight",
		Fields: gql.Fields{
			"id": &gql.Field{
				Type: gql.String,
			},
			"from": &gql.Field{
				Type: gql.String,
			},
			"to": &gql.Field{
				Type: gql.String,
			},
			"passenger_ids": &gql.Field{
				Type: gql.NewList(gql.String),
			},
		},
	},
)
