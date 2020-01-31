package queries

import (
	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
	"github.com/siddontang/go/log"
)

func GetCateoriesQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.Catgr),
		Description: "Get categories list",
		Resolve: func(p graphql.ResolveParams) (ctgrs interface{}, err error) {
			ctgrs, err = types.SelectAllCtgr()
			if err != nil {
				log.Error("error getting all  categories ", err)
				return
			}
			return
		},
	}
}
