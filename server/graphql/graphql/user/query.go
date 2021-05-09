package user

import gql "github.com/graphql-go/graphql"

type UserQuery struct {
	QueryName string
	Field     *gql.Field
}

func GetAllUserQueryFields() []UserQuery {
	queryField := newProductQueryField()
	return []UserQuery{
		{
			QueryName: "user",
			Field:     queryField.user(),
		},
	}
}
