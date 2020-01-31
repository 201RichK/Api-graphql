package types

import (
	"github.com/201RichK/graphql/db"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"

	"time"
)

type ActualiteMotCle struct {
	ID          int64      `json:"id"`
	Mot         string     `json:"mot" `
	Statut      bool       `json:"statut"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	CategorieID int64      `json:"categorie"`
}

func (t *ActualiteMotCle) TableName() string {
	return "actualite_mot_cle"
}

func init() {
	Db, err := db.Conn()
	defer Db.Close()
	if err != nil {
		log.Error(err)
	}
	Db.AutoMigrate(ActualiteMotCle{})
}

var MotCle = graphql.NewObject(graphql.ObjectConfig{
	Name: "motCle",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"mot":         &graphql.Field{Type: graphql.String},
		"statut":      &graphql.Field{Type: graphql.Boolean},
		"createdAt":   &graphql.Field{Type: graphql.DateTime},
		"updatedAt":   &graphql.Field{Type: graphql.DateTime},
		"categorieId": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
	},
})

// AddActualiteMotCle insert a new ActualiteMotCle into database and returns
// last inserted Id on success.
func AddActualiteMotCle(m *ActualiteMotCle) error {
	return db.Db.Create(m).Error
}

//SelectActualiteMotCle Select all ActualiteMotCle into the database and returns
//[]ActualiteArticel
func SelectAllMotCle() (motCles []*ActualiteMotCle, err error) {
	db, err := db.Conn()
	if err != nil {
		log.Warn(err)
	}
	defer db.Close()
	//var motCles []*ActualiteMotCle
	err = db.Model(&ActualiteMotCle{}).Find(&motCles).Error
	if err != nil {
		log.Error("Error quering table actualite_mot_cle", err)
	}

	return
}

//Select mot  where id = ?
func SelectMocleId(id int64) (mots []*ActualiteMotCle, err error) {
	db, err := db.Conn()
	if err != nil {
		log.Warn("Select mot cle by Id error ", err)
	}
	defer db.Close()

	err = db.Model(&ActualiteMotCle{}).Where("categorie_id = ?", id).Find(&mots).Error
	if err != nil {
		log.Errorf("SelectMocleId error ", id, err)
		return
	}
	return
}
