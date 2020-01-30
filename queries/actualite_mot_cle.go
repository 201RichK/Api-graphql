package queries

import (
	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
)

// GetUserQuery returns the queries available against user type.
func GetMotCleQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.MotCle),
		Resolve: func(param graphql.ResolveParams) (interface{}, error) {
			motCle, err := types.SelectAllMotCle()
			if err != nil {
				return nil, nil
			}
			return &motCle, nil
		},
	}
}
