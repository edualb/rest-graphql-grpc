package flight

import gql "github.com/graphql-go/graphql"

var FlightType = gql.NewObject(
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
