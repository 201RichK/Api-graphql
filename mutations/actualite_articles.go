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

//mutation for update actualite mot cle
func UpdateArticles() *graphql.Field {
	return &graphql.Field{
		Type:        types.Article,
		Description: "Update an articles ",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"titre": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"image": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"article_text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"lien": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"intro_text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"date": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"statut": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"categorieId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			artcicles := &types.ActualiteArticle{
				ID:           p.Args["id"].(int),
				Titre:        p.Args["titre"].(string),
				Image:        p.Args["image"].(string),
				ArticleTexte: p.Args["article_text"].(string),
				Lien:         p.Args["lien"].(string),
				IntroTexte:   p.Args["intro_text"].(string),
				Date:         p.Args["date"].(string),
				Statut:       p.Args["statut"].(bool),
				CategorieID:  p.Args["categorieId"].(int),
			}

			err := types.UpdateArticles(artcicles)
			if err != nil {
				return nil, err
			}

			return artcicles, nil
		},
	}
}
