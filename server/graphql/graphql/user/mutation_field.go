package user

import (
	gql "github.com/graphql-go/graphql"
)

type UserMutationField interface {
	createUser() *gql.Field
	updateUser() *gql.Field
	deleteUser() *gql.Field
}

type UserMutationFieldImpl struct{}

func newUserMutationField() UserMutationField {
	return &UserMutationFieldImpl{}
}

func (p *UserMutationFieldImpl) createUser() *gql.Field {
	return &gql.Field{
		Type:        userType,
		Description: "Create new user",
		Args: gql.FieldConfigArgument{
			"name": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"flight_ids": &gql.ArgumentConfig{
				Type: gql.NewList(gql.String),
			},
		},
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			// Access Mongo to create
			p := User{
				ID:        "int64(rand.Intn(100000))", // generate random ID
				Name:      params.Args["name"].(string),
				FlightIDs: params.Args["flight_ids"].([]string),
			}
			return p, nil
		},
	}
}

func (p *UserMutationFieldImpl) updateUser() *gql.Field {
	return &gql.Field{
		Type:        userType,
		Description: "Create new user",
		Args: gql.FieldConfigArgument{
			"name": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"flight_ids": &gql.ArgumentConfig{
				Type: gql.NewList(gql.String),
			},
		},
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			// Access Mongo to create
			p := User{
				ID:        "int64(rand.Intn(100000))", // generate random ID
				Name:      params.Args["name"].(string),
				FlightIDs: params.Args["flight_ids"].([]string),
			}
			return p, nil
		},
	}
}

func (p *UserMutationFieldImpl) deleteUser() *gql.Field {
	return &gql.Field{
		Type:        userType,
		Description: "Create new user",
		Args: gql.FieldConfigArgument{
			"name": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"flight_ids": &gql.ArgumentConfig{
				Type: gql.NewList(gql.String),
			},
		},
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			// Access Mongo to create
			p := User{
				ID:        "int64(rand.Intn(100000))", // generate random ID
				Name:      params.Args["name"].(string),
				FlightIDs: params.Args["flight_ids"].([]string),
			}
			return p, nil
		},
	}
}
