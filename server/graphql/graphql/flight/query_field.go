package flight

import gql "github.com/graphql-go/graphql"

type FlightQueryField interface {
	flight() *gql.Field
	flights() *gql.Field
}

type FlightQueryFieldImpl struct{}

func newFlightQueryField() FlightQueryField {
	return &FlightQueryFieldImpl{}
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
			_, ok := p.Args["id"].(string)
			if ok {
				// get user mongo by id
				// 	// Find product
				// 	for _, p := range Products {
				// 		if int(p.ID) == id {
				// 			return p, nil
				// 		}
				// 	}
			}
			return nil, nil
		},
	}
}

func (u *FlightQueryFieldImpl) flights() *gql.Field {
	return &gql.Field{
		Type:        gql.NewList(userType),
		Description: "Get product list",
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			// get all users
			return nil, nil
		},
	}
}
