package mutation

import (
	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
	"time"
)

func CreateCategorie() *graphql.Field {
	return &graphql.Field{
		Type:        types.Catgr,
		Description: "Create actualite categories",
		Args: graphql.FieldConfigArgument{
			"Nom": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"Statut": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			catgr := types.ActualiteCategorie{
				Nom:       params.Args["Nom"].(string),
				Statut:    params.Args["statut"].(bool),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			err := types.AddActualiteCtgr(&catgr)
			if err != nil {
				return nil, err
			}

			return catgr, nil
		},
	}
}
