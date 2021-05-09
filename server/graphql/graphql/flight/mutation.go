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
			MutationName: "createFlight",
			Field:        queryField.createFlight(),
		},
		{
			MutationName: "updateFlight",
			Field:        queryField.updateFlight(),
		},
	}
}
