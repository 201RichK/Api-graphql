package queries

import (
	"github.com/201RichK/graphql/types"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

func GetCateoriesQueries() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.Catgr),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "include Id to get by Id",
			},
		},
		Description: "Get categories list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if id, okay := p.Args["id"].(int); okay {
				ctgr, err := types.SelectAllctgrId(id, "id")
				if err != nil {
					return nil, err
				}
				return ctgr, nil
			}

			ctgrs, err := types.SelectAllCtgr()
			if err != nil {
				log.Error("error getting all  categories ", err)
				return nil, err
			}
			return ctgrs, nil
		},
	}
}
