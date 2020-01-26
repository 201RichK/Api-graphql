package types

import (
	"github.com/201RichK/graphql/db"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type ActualiteCategorie struct {
	gorm.Model       `json:"model"`
	Nom              string `json:"nom"`
	Statut           bool   `json:"statut"`
	ActualiteMotCle  []ActualiteMotCle
	ActualiteArticle []ActualiteArticle
}

func (t *ActualiteCategorie) TableName() string {
	return "actualite_categorie"
}

func init() {
	db.Db.AutoMigrate(ActualiteCategorie{})
}

//usertype is use by the GRAPHQL API to specify which field can be access
var actualiteCategorie = graphql.NewObject(graphql.ObjectConfig{
	Name: "Actualite_Categorie",
	Fields: graphql.Fields{
		"nom":    &graphql.Field{Type: graphql.String},
		"statut": &graphql.Field{Type: graphql.Boolean},
	},
})
