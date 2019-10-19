package mutation

import (
	"github.com/201RichK/graphql/graphql-test/types"
	"github.com/graphql-go/graphql"
)

func GetRootField() graphql.Fields {
	return graphql.Fields{
		"createUser": GetCreateMutation(),
	}
}

// GetCreateUserMutation creates a new user and returns it.
func GetCreateMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			user := &types.Users{
				Firstname: params.Args["firsname"].(string),
				Lastname:  params.Args["lastname"].(string),
			}

			//Add user to database here

			return user, nil
		},
	}
}
