package flight

import (
	gql "github.com/graphql-go/graphql"
)

type FlightMutationField interface {
	createFlight() *gql.Field
	updateFlight() *gql.Field
	deleteFlight() *gql.Field
}

type FlightMutationFieldImpl struct{}

func newFlightMutationField() FlightMutationField {
	return &FlightMutationFieldImpl{}
}

func (p *FlightMutationFieldImpl) createFlight() *gql.Field {
	return &gql.Field{
		Type:        userType,
		Description: "Create new user",
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
			// Access Mongo to create
			p := Flight{
				ID:           "int64(rand.Intn(100000))", // generate random ID
				From:         params.Args["from"].(string),
				PassengerIDs: params.Args["passenger_ids"].([]string),
			}
			return p, nil
		},
	}
}

func (p *FlightMutationFieldImpl) updateFlight() *gql.Field {
	return &gql.Field{
		Type:        userType,
		Description: "Create new user",
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
			// Access Mongo to create
			p := Flight{
				ID:           "int64(rand.Intn(100000))", // generate random ID
				From:         params.Args["from"].(string),
				PassengerIDs: params.Args["passenger_ids"].([]string),
			}
			return p, nil
		},
	}
}

func (p *FlightMutationFieldImpl) deleteFlight() *gql.Field {
	return &gql.Field{
		Type:        userType,
		Description: "Create new user",
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
			// Access Mongo to create
			p := Flight{
				ID:           "int64(rand.Intn(100000))", // generate random ID
				From:         params.Args["from"].(string),
				PassengerIDs: params.Args["passenger_ids"].([]string),
			}
			return p, nil
		},
	}
}
