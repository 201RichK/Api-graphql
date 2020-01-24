package types

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

type ActualiteMotCle struct {
	Id          int       `json:"id"`
	Mot         string    `json:"mot"`
	Statut      bool      `json:"statut"`
	DateUpd     time.Time `json:"date_upd"`
	DateAdd     time.Time `json:"date_add"`
	CategorieId int64     `json:"categorie_id"`
}

func (t *ActualiteMotCle) TableName() string {
	return "actualite_article"
}

func init() {
	orm.RegisterModel(new(ActualiteMotCle))
}

var actualiteMotCle = graphql.NewObject(graphql.ObjectConfig{
	Name: "Actualite_mot_cle",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.Int},
		"mot":          &graphql.Field{Type: graphql.String},
		"statut":       &graphql.Field{Type: graphql.Boolean},
		"date_add":     &graphql.Field{Type: graphql.DateTime},
		"data_upd":     &graphql.Field{Type: graphql.DateTime},
		"categorie_id": &graphql.Field{Type: graphql.Int},
	},
})
