package user

import (
	"github.com/edualb/rest-graphql-grpc/server/graphql/models"
	gql "github.com/graphql-go/graphql"
)

type UserQueryField interface {
	user() *gql.Field
}

type UserQueryFieldImpl struct {
	db models.UsersDB
}

func newProductQueryField() UserQueryField {
	return &UserQueryFieldImpl{
		db: models.NewUsersDB(),
	}
}

func (u *UserQueryFieldImpl) user() *gql.Field {
	return &gql.Field{
		Type:        usersType,
		Description: "Get product by id",
		Args: gql.FieldConfigArgument{
			"id": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(string)
			if ok {
				user, err := u.db.FindUserByID(id)
				if err != nil {
					return nil, err
				}
				return user, nil
			}
			return nil, nil
		},
	}
}
