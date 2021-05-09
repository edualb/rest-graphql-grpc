package flight

import (
	"github.com/edualb/rest-graphql-grpc/server/graphql/models"
	gql "github.com/graphql-go/graphql"
)

type FlightQueryField interface {
	flight() *gql.Field
}

type FlightQueryFieldImpl struct {
	db models.FlightsDB
}

func newFlightQueryField() FlightQueryField {
	return &FlightQueryFieldImpl{
		db: models.NewFlightsDB(),
	}
}

func (u *FlightQueryFieldImpl) flight() *gql.Field {
	return &gql.Field{
		Type:        userType,
		Description: "Get product by id",
		Args: gql.FieldConfigArgument{
			"id": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(string)
			if ok {
				flight, err := u.db.FindFlightByID(id)

				if err != nil {
					return nil, err
				}

				return flight, nil
			}
			return nil, nil
		},
	}
}
