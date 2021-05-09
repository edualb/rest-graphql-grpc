package flight

import gql "github.com/graphql-go/graphql"

type FlightMutation struct {
	MutationName string
	Field        *gql.Field
}

func GetAllFlightMutationFields() []FlightMutation {
	queryField := newFlightMutationField()
	return []FlightMutation{
		{
			MutationName: "createUser",
			Field:        queryField.createFlight(),
		},
		{
			MutationName: "updateUser",
			Field:        queryField.updateFlight(),
		},
		{
			MutationName: "deleteUser",
			Field:        queryField.deleteFlight(),
		},
	}
}
