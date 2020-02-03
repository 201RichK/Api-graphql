package mutation

import (
	"time"

	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

// mutation for creating actualite mot cle
func CreateMocle() *graphql.Field {
	return &graphql.Field{
		Type:        types.MotCle,
		Description: "create new mot clé for the actualite table",
		Args: graphql.FieldConfigArgument{
			"mot": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"statut": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"created_at": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"categorie_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			mtcle := &types.ActualiteMotCle{
				Mot:         params.Args["mot"].(string),
				Statut:      params.Args["statut"].(bool),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				CategorieID: params.Args["categorie_id"].(int),
			}
			err := types.AddActualiteMotCle(mtcle)
			if err != nil {
				log.Error("CreateMocle error ", err)
				return nil, err
			}

			return mtcle, nil
		},
	}
}

//mutation for update actualite mot cle
func UpdateMotCle() *graphql.Field {
	return &graphql.Field{
		Type:        types.MotCle,
		Description: "Update a mot cle for actualite table",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
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
			motCle := &types.ActualiteMotCle{
				ID:          p.Args["id"].(int),
				Statut:      p.Args["statut"].(bool),
				Mot:         p.Args["mot"].(string),
				UpdatedAt:   time.Now(),
				CategorieID: p.Args["categorie_id"].(int),
			}

			err := types.UpdateMotCle(motCle)
			if err != nil {
				return nil, err
			}

			return motCle, nil
		},
	}
}
