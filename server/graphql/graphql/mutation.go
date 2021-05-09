package graphql

import (
	"github.com/edualb/rest-graphql-grpc/server/graphql/graphql/flight"
	"github.com/edualb/rest-graphql-grpc/server/graphql/graphql/user"
	gql "github.com/graphql-go/graphql"
)

var mutationFields gql.Fields

func init() {
	mutationFields = make(gql.Fields)

	for _, userMutation := range user.GetAllUserMutationFields() {
		mutationFields[userMutation.MutationName] = userMutation.Field
	}

	for _, flightMutation := range flight.GetAllFlightMutationFields() {
		mutationFields[flightMutation.MutationName] = flightMutation.Field
	}
}

func mutation() *gql.Object {
	return gql.NewObject(
		gql.ObjectConfig{
			Name:   "Mutation",
			Fields: mutationFields,
		},
	)
}
