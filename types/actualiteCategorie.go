package types

import (
	"time"

	"github.com/201RichK/graphql/db"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

type ActualiteCategorie struct {
	ID               int64      `json:"id"`
	Nom              string     `json:"nom"`
	Statut           bool       `json:"statut"`
	CreatedAt        *time.Time `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
	ActualiteMotCle  []ActualiteMotCle
	ActualiteArticle []ActualiteArticle
}

func (t *ActualiteCategorie) TableName() string {
	return "actualite_categorie"
}

func init() {
	Db, err := db.Conn()
	defer Db.Close()
	if err != nil {
		log.Error(err)
	}
	Db.AutoMigrate(ActualiteCategorie{})
}

//usertype is use by the GRAPHQL API to specify which field can be access
var Catgr = graphql.NewObject(graphql.ObjectConfig{
	Name: "Actualite_Categorie",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"nom":       &graphql.Field{Type: graphql.String},
		"statut":    &graphql.Field{Type: graphql.Boolean},
		"createdAt": &graphql.Field{Type: graphql.DateTime},
		"updatedAt": &graphql.Field{Type: graphql.DateTime},
		"actualiteMotCle": &graphql.Field{
			Type: graphql.NewList(MotCle),
			Resolve: func(p graphql.ResolveParams) (Mocles interface{}, err error) {
				id := p.Source.(*ActualiteCategorie).ID

				Mocles, err = SelectMocleId(id)
				if err != nil {
					return
				}
				return
			},
		},
		"actualiteArticle": &graphql.Field{
			Type: graphql.NewList(Article),
			Resolve: func(p graphql.ResolveParams) (articles interface{}, err error) {
				id := p.Source.(*ActualiteCategorie).ID

				articles, err = SelectAllArtId(id)
				if err != nil {
					return
				}

				return
			},
		},
	},
})

func SelectAllCtgr() (actCtgr []*ActualiteCategorie, err error) {
	db, err := db.Conn()
	if err != nil {
		log.Warn(err)
	}
	defer db.Close()
	//var motCles []*ActualiteMotCle
	err = db.Model(&ActualiteCategorie{}).Find(&actCtgr).Error
	if err != nil {
		log.Error("Error quering table actualite_mot_cle", err)
	}

	return
}

func SelectAllctgrId() (ctgr []*ActualiteCategorie, err error) {
	db, err := db.Conn()
	if err != nil {
		log.Error("error select all categories where id ==> ", err)
	}

	err = db.Model(ActualiteCategorie{}).Where("categorie_id = ?").Find(&ctgr).Error
	if err != nil {
		log.Error("SelectAllctgrId request error ", err)

		return
	}

	return
}
