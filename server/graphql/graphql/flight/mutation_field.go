package flight

import (
	"github.com/edualb/rest-graphql-grpc/server/graphql/models"
	gql "github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightMutationField interface {
	createFlight() *gql.Field
	updateFlight() *gql.Field
}

type FlightMutationFieldImpl struct {
	db models.FlightsDB
}

func newFlightMutationField() FlightMutationField {
	return &FlightMutationFieldImpl{
		db: models.NewFlightsDB(),
	}
}

func (p *FlightMutationFieldImpl) createFlight() *gql.Field {
	return &gql.Field{
		Type:        userType,
		Description: "Create new flight",
		Args: gql.FieldConfigArgument{
			"from": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"to": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"passenger_ids": &gql.ArgumentConfig{
				Type: gql.NewList(gql.String),
			},
		},
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			flight := models.Flights{}

			passengerIDs, passengerIDsOK := params.Args["flight_ids"].([]string)

			if passengerIDsOK {
				ids := []primitive.ObjectID{}
				for _, id := range passengerIDs {
					primitiveID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						return nil, err
					}
					ids = append(ids, primitiveID)
				}
				flight.PassengerIDs = ids
			}

			from, fromOK := params.Args["from"].(string)
			to, toOK := params.Args["to"].(string)

			if fromOK && toOK {
				flight.From = from
				flight.To = to

				err := p.db.CreateFlight(flight)

				if err != nil {
					return nil, err
				}
				return flight, err
			}
			return flight, nil
		},
	}
}

func (p *FlightMutationFieldImpl) updateFlight() *gql.Field {
	return &gql.Field{
		Type:        userType,
		Description: "Create new user",
		Args: gql.FieldConfigArgument{
			"id": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"from": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"to": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"passenger_ids": &gql.ArgumentConfig{
				Type: gql.NewList(gql.String),
			},
		},
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			id, idOK := params.Args["id"].(string)

			if idOK {
				flight := models.Flights{}

				passengerIDs, passengerIDsOK := params.Args["flight_ids"].([]string)

				if passengerIDsOK {
					ids := []primitive.ObjectID{}
					for _, id := range passengerIDs {
						primitiveID, err := primitive.ObjectIDFromHex(id)
						if err != nil {
							return nil, err
						}
						ids = append(ids, primitiveID)
					}
					flight.PassengerIDs = ids
				}

				from, fromOK := params.Args["from"].(string)
				to, toOK := params.Args["to"].(string)

				if fromOK && toOK {
					flight.From = from
					flight.To = to

					err := p.db.UpdateFlightByID(id, flight)

					if err != nil {
						return nil, err
					}
					return flight, err
				}
				return flight, nil
			}
			return nil, nil
		},
	}
}
