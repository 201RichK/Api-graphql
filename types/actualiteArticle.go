package types

import (
	"fmt"
	"github.com/201RichK/graphql/db"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"time"
)

type ActualiteArticle struct {
	ID           int       `json:"id" gorm:"AUTO_INCREMENT"`
	Titre        string    `json:"titre"`
	Image        string    `json:"image"`
	ArticleTexte string    `json:"article_text"`
	Lien         string    `json:"lien"`
	IntroTexte   string    `json:"intro_text"`
	Date         string    `json:"date"`
	Statut       bool      `json:"statut"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CategorieID  int       `json:"categorieId"`
}

func (t *ActualiteArticle) TableName() string {
	return "actualite_article"
}

func init() {
	Db, err := db.Conn()
	defer Db.Close()
	if err != nil {
		log.Error(err)
	}
	Db.AutoMigrate(ActualiteArticle{})
}

//actualiteArticle type use to specify which field can access
var Article = graphql.NewObject(graphql.ObjectConfig{
	Name: "articles",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.Int},
		"titre":        &graphql.Field{Type: graphql.String},
		"image":        &graphql.Field{Type: graphql.String},
		"article_text": &graphql.Field{Type: graphql.String},
		"lien":         &graphql.Field{Type: graphql.String},
		"intro_text":   &graphql.Field{Type: graphql.String},
		"date":         &graphql.Field{Type: graphql.String},
		"statut":       &graphql.Field{Type: graphql.Boolean},
		"categorieId":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
	},
})

// AddActualiteMotCle insert a new ActualiteMotCle into database and returns
// last inserted Id on success.
func AddActualiteArticles(m *ActualiteArticle) error {
	db, err := db.Conn()
	if err != nil {
		log.Error("AddActualiteArticle Can't connect to database ", err)
	}
	defer db.Close()

	return db.Where(ActualiteArticle{Titre: m.Titre}).Attrs(ActualiteArticle{Titre: m.Titre}).FirstOrCreate(m).Error
}

func SelectAllArticle() (articles []*ActualiteArticle, err error) {
	db, err := db.Conn()
	if err != nil {
		log.Error("can't connect to database ", err)
		return
	}
	defer db.Close()
	//var articles []*ActualiteArticle

	err = db.Model(&ActualiteArticle{}).Find(&articles).Error
	if err != nil {
		log.Error("error query Actualite_article ", err)
		return
	}
	return
}

func SelectAllArtId(id int, table string) (artcicle []*ActualiteArticle, err error) {
	db, err := db.Conn()
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	err = db.Model(&ActualiteArticle{}).Where(fmt.Sprintf("%s = ?", table), id).Find(&artcicle).Error
	if err != nil {
		log.Error("SelectAllArtId error ", err)
		return
	}
	return
}

//Update actualite Articles
func UpdateArticles(m *ActualiteArticle) (err error) {
	db, err := db.Conn()
	if err != nil {
		return
	}
	defer db.Close()
	if err = db.Model(&ActualiteArticle{}).Where("id = ?", m.ID).Take(&ActualiteArticle{}).Save(m).Error; err != nil {
		return
	}

	return
}
