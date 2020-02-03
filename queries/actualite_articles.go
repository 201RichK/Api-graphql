package queries

import (
	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

func GetArticleQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.Article),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Id to get all artticle by id",
			},
			"categorieId": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Categorie Id to get all articles by categorieId",
			},
		},
		Description: "Get articles list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			if id, ok := p.Args["id"].(int); ok {
				atcl, err := types.SelectAllArtId(id, "id")
				if err != nil {
					return nil, err
				}
				return atcl, nil
			}

			if ctgrID, ok := p.Args["categorieId"].(int); ok {
				artcl, err := types.SelectAllArtId(ctgrID, "categorie_id")
				if err != nil {
					return nil, err
				}
				return artcl, nil
			}

			articles, err := types.SelectAllArticle()
			if err != nil {
				logrus.Info("error select all articles ", err)
				return nil, nil
			}
			return articles, nil
		},
	}
}
