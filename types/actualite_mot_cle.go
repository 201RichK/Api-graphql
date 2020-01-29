package types

import (
	"github.com/201RichK/graphql/db"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type ActualiteMotCle struct {
	gorm.Model  `json:"model"`
	Mot         string `json:"mot" `
	Statut      bool   `json:"statut"`
	CategorieID int64  `json:"categorie"`
}

func (t *ActualiteMotCle) TableName() string {
	return "actualite_mot_cle"
}

func init() {
	db.Db.AutoMigrate(ActualiteMotCle{})
}

var MotCle = graphql.NewObject(graphql.ObjectConfig{
	Name: "Actualite_mot_cle",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.Int},
		"mot":    &graphql.Field{Type: graphql.String},
		"statut": &graphql.Field{Type: graphql.Boolean},
		"categorie": &graphql.Field{
			Type: graphql.NewList(actualiteCategorie),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var categorie ActualiteCategorie

				return categorie, nil
			},
		},
	},
})

// AddActualiteMotCle insert a new ActualiteMotCle into database and returns
// last inserted Id on success.
func AddActualiteMotCle(m *ActualiteMotCle) error {
	return db.Db.Create(m).Error
}

//SelectActualiteMotCle Select all ActualiteMotCle into the database and returns
//[]ActualiteArticel
func SelectAllMotCle() ([]*ActualiteMotCle, error) {
	var motCle []*ActualiteMotCle
	err := db.Db.Find(ActualiteMotCle{}).Scan(motCle).Error  
	if err != nil {
		return nil, err
	}

	return motCle, nil
}
