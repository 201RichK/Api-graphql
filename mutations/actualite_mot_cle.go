package mutation

import (
	"fmt"
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
			todo := &types.ActualiteMotCle{
				Mot:         params.Args["mot"].(string),
				Statut:      params.Args["statut"].(bool),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				CategorieID: params.Args["categorie_id"].(int),
			}
			err := types.AddActualiteMotCle(todo)
			log.Println(err)
			if err != nil {
				log.Error("CreateMocle error ", err)
				return nil, err
			}

			return todo, nil
		},
	}
}

//mutation for update actualite mot cle
func UpdateMotCle() *graphql.Field {
	return &graphql.Field{
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
	}
}
