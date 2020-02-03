package mutation

import (
	"time"

	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
)

func CreateArticle() *graphql.Field {
	return &graphql.Field{
		Type:        types.Article,
		Description: "Create a new aricles",
		Args: graphql.FieldConfigArgument{
			"Titre": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"Image": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"ArticleText": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"Lien": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"IntroText": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"Date": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"Statut": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"categorieID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			atcl := types.ActualiteArticle{
				Titre:        params.Args["titre"].(string),
				Image:        params.Args["Image"].(string),
				ArticleTexte: params.Args["ArticleText"].(string),
				Lien:         params.Args["Lien"].(string),
				IntroTexte:   params.Args["IntroText"].(string),
				Date:         params.Args["Date"].(string),
				Statut:       params.Args["Statut"].(bool),
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				CategorieID:  params.Args["categorieID"].(int),
			}

			err := types.AddActualiteArticles(&atcl)
			if err != nil {
				return nil, err
			}
			return atcl, nil
		},
	}
}
