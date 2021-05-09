package user

import gql "github.com/graphql-go/graphql"

type UserQueryField interface {
	user() *gql.Field
	users() *gql.Field
}

type UserQueryFieldImpl struct{}

func newProductQueryField() UserQueryField {
	return &UserQueryFieldImpl{}
}

func (u *UserQueryFieldImpl) user() *gql.Field {
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

func (u *UserQueryFieldImpl) users() *gql.Field {
	return &gql.Field{
		Type:        gql.NewList(userType),
		Description: "Get all users",
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			// get all users
			return nil, nil
		},
	}
}
