package user

import gql "github.com/graphql-go/graphql"

type UserMutation struct {
	MutationName string
	Field        *gql.Field
}

func GetAllUserMutationFields() []UserMutation {
	queryField := newUserMutationField()
	return []UserMutation{
		{
			MutationName: "createUser",
			Field:        queryField.createUser(),
		},
		{
			MutationName: "updateUser",
			Field:        queryField.updateUser(),
		},
		{
			MutationName: "deleteUser",
			Field:        queryField.deleteUser(),
		},
	}
}
