package mutation

import (
	"fmt"

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

		// Create new actualite_mot_cle
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
				"categorie_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				err := types.AddActualiteMotCle(&types.ActualiteMotCle{
					Mot:         params.Args["mot"].(string),
					Statut:      params.Args["statut"].(bool),
					CategorieID: params.Args["categorie_id"].(int64),
				})
				if err != nil {
					return nil, err
				}

				return "Okay", nil
			},
		},
		"update": &graphql.Field{
			Type:        types.MotCle,
			Description: "Update a mot cle for actualite table",
			Args: graphql.FieldConfigArgument{
				"mot": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"statut": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
				"categorie_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				mot, motOk := p.Args["mot"].(string)
				statut, statutOk := p.Args["statut"].(bool)
				ctgrID, ctgrIDOk := p.Args["categorie_id"].(int)

				fmt.Println(mot, motOk, statut, statutOk, ctgrID, ctgrIDOk)

				return nil, nil
			},
		},
	}
}
