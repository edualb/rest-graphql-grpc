package graphql

import (
	"github.com/edualb/rest-graphql-grpc/server/graphql/graphql/flight"
	"github.com/edualb/rest-graphql-grpc/server/graphql/graphql/user"
	gql "github.com/graphql-go/graphql"
)

var queryFields gql.Fields

func init() {
	queryFields = make(gql.Fields)

	for _, userQuery := range user.GetAllUserQueryFields() {
		queryFields[userQuery.QueryName] = userQuery.Field
	}

	for _, flightQuery := range flight.GetAllFlightQueryFields() {
		queryFields[flightQuery.QueryName] = flightQuery.Field
	}
}

func query() *gql.Object {

	return gql.NewObject(
		gql.ObjectConfig{
			Name:   "Query",
			Fields: queryFields,
		})
}
