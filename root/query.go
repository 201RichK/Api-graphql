package root

import (
	"github.com/201RichK/graphql/queries"
	"github.com/graphql-go/graphql"
)

//GeRootFiels return all  the available queries
func GetMotCleField() graphql.Fields {
	return graphql.Fields{
		//Mot cle
		"motCles": queries.GetMotCleQuery(),

		//Articles
		"articles": queries.GetArticleQuery(),

		//Catgeories
		"categories": queries.GetCateoriesQueries(),
	}
}
