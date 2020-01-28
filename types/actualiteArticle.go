package types

import (
	"github.com/201RichK/graphql/db"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type ActualiteArticle struct {
	gorm.Model   `json:"model"`
	Titre        string `json:"titre"`
	Image        string `json:"image"`
	ArticleTexte string `json"article_text"`
	Lien         string `json:"lien"`
	IntroTexte   string `json:"intro_texte"`
	Date         string `json:"date"`
	Statut       bool   `json:"statut"`
	CategorieID  int64  `json:"categorie"`
}

func (t *ActualiteArticle) TableName() string {
	return "actualite_article"
}

func init() {
	db.Db.AutoMigrate(ActualiteArticle{})
}

//actualiteArticle type use to specify which field can access
var actualiteArticle = graphql.NewObject(graphql.ObjectConfig{
	Name: "Actualite_Article",
	Fields: graphql.Fields{
		"titre":        &graphql.Field{Type: graphql.String},
		"image":        &graphql.Field{Type: graphql.String},
		"article_text": &graphql.Field{Type: graphql.String},
		"lien":         &graphql.Field{Type: graphql.String},
		"intro_text":   &graphql.Field{Type: graphql.String},
		"date":         &graphql.Field{Type: graphql.String},
		"statut":       &graphql.Field{Type: graphql.Boolean},
		"categorie": &graphql.Field{
			Type: graphql.NewList(actualiteCategorie),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var categorie ActualiteCategorie

				return categorie, nil
			},
		},
	},
})
