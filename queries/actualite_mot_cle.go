package queries

import (
	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
)

// GetUserQuery returns the queries available against user type.
func GetMotCleQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.MotCle),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Include ID to get by id",
			},
		},
		Description: "Get mot cle list",
		Resolve: func(param graphql.ResolveParams) (interface{}, error) {

			if id, ok := param.Args["id"].(int); ok {
				ctgr, err := types.SelectMocleId(id, "id")
				if err != nil {
					return nil, err
				}
				return ctgr, nil
			}

			motCle, err := types.SelectAllMotCle()
			if err != nil {
				return nil, nil
			}
			return motCle, nil
		},
	}
}
