package root

import (
	mutation "github.com/201RichK/graphql/mutations"
	"github.com/graphql-go/graphql"
)

//GetMutationField return all  the available mutation fields
func GetMutationField() graphql.Fields {
	return graphql.Fields{
		"create": mutation.CreateMocle(),
		"update": mutation.UpdateMotCle(),
	}
}
