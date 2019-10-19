package queries

import (
	"github.com/201RichK/graphql/graphql-test/types"
	"github.com/graphql-go/graphql"
)

//GeRootFiels return all  the available queries
func GetRootField() graphql.Fields {
	return graphql.Fields{
		"user": GetUserQuery(),
	}
}

// GetUserQuery returns the queries available against user type.
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(param graphql.ResolveParams) (interface{}, error) {
			var users []types.Users

			//database logique

			return users, nil
		},
	}
}
