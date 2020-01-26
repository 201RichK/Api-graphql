package mutation

import (
	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
)

// GetRootFields returns all the available mutations.
func GetRootField() graphql.Fields {
	return GetCreateMutation()
}

// GetCreateUserMutation creates a new user and returns it.
func GetCreateMutation() graphql.Fields {
	return graphql.Fields{
		/*
			Create new actualite_mot_cle
		*/
		"create": &graphql.Field{
			Type:        types.MotCle,
			Description: "create new mot clé for the actualite table",
			Args: graphql.FieldConfigArgument{
				"mot": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"statut": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
				"date_add": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"data_upd": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"categorie_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) error {

				err := types.AddActualiteMotCle(&types.ActualiteMotCle{
					Mot:         params.Args["mot"].(string),
					Statut:      params.Args["statut"].(bool),
					CategorieID: params.Args["categorie_id"].(int64),
				})
				if err != nil {
					return err
				}

				return nil
			},
		},
	}
}
