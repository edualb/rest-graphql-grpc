package flight

import gql "github.com/graphql-go/graphql"

type UserQuery struct {
	QueryName string
	Field     *gql.Field
}

func GetAllFlightQueryFields() []UserQuery {
	queryField := newFlightQueryField()
	return []UserQuery{
		{
			QueryName: "flight",
			Field:     queryField.flight(),
		},
		{
			QueryName: "flights",
			Field:     queryField.flights(),
		},
	}
}
