package types

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

type ActualiteCategorie struct {
	Id      int       `json:"id"`
	Nom     string    `json:"nom"`
	Statut  bool      `json:"statut"`
	DateAdd time.Time `json:"date_add"`
	DateUpd time.Time `json:"date_upd"`
}


func (t *ActualiteCategorie) TableName() string {
	return "actualite_categorie"
}

func init() {
	orm.RegisterModel(new(ActualiteCategorie))
}

//usertype is use by the GRAPHQL API to specify which field can be access
var actualiteCategorie = graphql.NewObject(graphql.ObjectConfig{
	Name: "Actualite_Categorie",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.Int},
		"nom":      &graphql.Field{Type: graphql.String},
		"statut":   &graphql.Field{Type: graphql.Boolean},
		"date_add": &graphql.Field{Type: graphql.DateTime},
		"date_upd": &graphql.Field{Type: graphql.DateTime},
	},
})
