package queries

import (
	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
)

//GeRootFiels return all  the available queries
func GetRootField() graphql.Fields {
	return graphql.Fields{
		"actualite_mot_cle": GetMotCleQuery(),
	}
}

// GetUserQuery returns the queries available against user type.
func GetMotCleQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.MotCle),
		Description: "Get mot cle list",
		Resolve: func(param graphql.ResolveParams) (interface{}, error) {
			var motCle []*types.ActualiteMotCle

			motCle, err := types.SelectAllMotCle()
			if err != nil {
				return nil, err
			}
			return motCle, nil
		},
	}
}
