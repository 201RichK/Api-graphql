package queries

import (
	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

func GetArticleQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.Article),
		Description: "Get articles list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			articles, err := types.SelectAllArticle()
			if err != nil {
				logrus.Info("error select all articles ", err)
				return nil, nil
			}
			return articles, nil
		},
	}
}
