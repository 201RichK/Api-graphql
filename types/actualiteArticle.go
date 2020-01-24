package types

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

type ActualiteArticle struct {
	Id           int                 `json:"id"`
	Titre        string              `json:"titre"`
	Image        string              `json:"image"`
	ArticleTexte string              `json"article_text"`
	Lien         string              `json:"lien"`
	IntroTexte   string              `json:"intro_texte"`
	Date         string              `json:"date"`
	Statut       bool                `json:"statut"`
	DateAdd      string              `json:"date_add"`
	DateUpd      time.Time           `json:"date_upd"`
	Categorie    *ActualiteCategorie `json:"categorie"`
}

func (t *ActualiteArticle) TableName() string {
	return "actualite_article"
}

func init() {
	orm.RegisterModel(new(ActualiteArticle))
}

//actualiteArticle type use to specify which field can access
var actualiteArticle = graphql.NewObject(graphql.ObjectConfig{
	Name: "Actualite_Article",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.Int},
		"titre":        &graphql.Field{Type: graphql.String},
		"image":        &graphql.Field{Type: graphql.String},
		"article_text": &graphql.Field{Type: graphql.String},
		"lien":         &graphql.Field{Type: graphql.String},
		"intro_text":   &graphql.Field{Type: graphql.String},
		"date":         &graphql.Field{Type: graphql.String},
		"statut":       &graphql.Field{Type: graphql.Boolean},
		"date_add":     &graphql.Field{Type: graphql.String},
		"date_upd":     &graphql.Field{Type: graphql.DateTime},
		"categorie": &graphql.Field{
			Type: graphql.NewList(actualiteCategorie),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var categorie ActualiteCategorie

				return categorie, nil
			},
		},
	},
})
