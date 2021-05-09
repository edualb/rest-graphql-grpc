package graphql

import (
	"fmt"

	gql "github.com/graphql-go/graphql"
)

type Schema struct{}

func NewSchema() *gql.Schema {
	schema, err := gql.NewSchema(
		gql.SchemaConfig{
			Query:    query(),
			Mutation: mutation(),
		},
	)
	if err != nil {
		panic(fmt.Sprintf("Graphql schema does not created: %s", err.Error()))
	}
	return &schema
}
