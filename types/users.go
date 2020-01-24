package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

type ActualiteCategorie struct {
	Id int `json:"id"`
	Nom string `json:"nom"`
	Statut bool `json:"statut"`
	DateAdd time.Time `json:"date_add"`
	DateUpd time.Time `json:"date_upd"`
}

type Users struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      []*Role
}

//user role definition
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//usertype is use bys the GRAPHQL API to specify which field can be access
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"firstname": &graphql.Field{Type: graphql.String},
		"lastname":  &graphql.Field{Type: graphql.String},
		"roles": &graphql.Field{
			Type: graphql.NewList(roleType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var role []Role

				// userID := params.Source.(User).ID
				// Implement logic to retrieve user associated roles from user id here.

				return role, nil
			},
		},
	},
})

var roleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String},
	},
})
