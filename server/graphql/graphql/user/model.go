package user

import (
	"github.com/edualb/rest-graphql-grpc/server/graphql/graphql/flight"
	"github.com/edualb/rest-graphql-grpc/server/graphql/models"
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
			"flights": &gql.Field{
				Type: gql.NewList(flight.FlightType),
				Resolve: func(params gql.ResolveParams) (interface{}, error) {

					db := models.NewFlightsDB()

					flights := []models.Flights{}

					users := params.Source.(*models.Users)

					for _, flight_id := range users.FlightIDs {
						flight, err := db.FindFlightByID(flight_id.Hex())
						if err != nil {
							return nil, err
						}
						flights = append(flights, *flight)
					}

					return flights, nil
				},
			},
		},
	},
)
