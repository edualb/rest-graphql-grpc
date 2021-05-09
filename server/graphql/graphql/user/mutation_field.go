package user

import (
	"github.com/edualb/rest-graphql-grpc/server/graphql/models"
	gql "github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMutationField interface {
	createUser() *gql.Field
	updateUser() *gql.Field
	deleteUser() *gql.Field
}

type UserMutationFieldImpl struct {
	db models.UsersDB
}

func newUserMutationField() UserMutationField {
	return &UserMutationFieldImpl{
		db: models.NewUsersDB(),
	}
}

func (p *UserMutationFieldImpl) createUser() *gql.Field {
	return &gql.Field{
		Type:        usersType,
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
			user := models.Users{}

			flightIDs, flightIDsOK := params.Args["flight_ids"].([]string)

			if flightIDsOK {
				ids := []primitive.ObjectID{}
				for _, id := range flightIDs {
					primitiveID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						return nil, err
					}
					ids = append(ids, primitiveID)
				}
				user.FlightIDs = ids
			}

			name, nameOk := params.Args["name"].(string)

			if nameOk {
				user.Name = name
				p.db.CreateUser(user)
				return user, nil
			}

			return user, nil
		},
	}
}

func (p *UserMutationFieldImpl) updateUser() *gql.Field {
	return &gql.Field{
		Type:        usersType,
		Description: "Create new user",
		Args: gql.FieldConfigArgument{
			"id": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"name": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"flight_ids": &gql.ArgumentConfig{
				Type: gql.NewList(gql.String),
			},
		},
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			user := models.Users{}

			id, idOK := params.Args["id"].(string)

			if idOK {

				flightIDs, flightIDsOK := params.Args["flight_ids"].([]string)

				if flightIDsOK {
					ids := []primitive.ObjectID{}
					for _, id := range flightIDs {
						primitiveID, err := primitive.ObjectIDFromHex(id)
						if err != nil {
							return nil, err
						}
						ids = append(ids, primitiveID)
					}
					user.FlightIDs = ids
				}

				name, nameOk := params.Args["name"].(string)

				if nameOk {
					user.Name = name
					err := p.db.UpdateUser(id, user)
					if err != nil {
						return nil, err
					}
				}
				return user, nil
			}

			return user, nil
		},
	}
}

func (p *UserMutationFieldImpl) deleteUser() *gql.Field {
	return &gql.Field{
		Type:        usersType,
		Description: "Create new user",
		Args: gql.FieldConfigArgument{
			"id": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
		},
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			id, idOK := params.Args["id"].(string)

			if idOK {
				err := p.db.DeleteUserByID(id)
				if err != nil {
					return nil, err
				}
			}

			return nil, nil
		},
	}
}
