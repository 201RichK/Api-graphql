package types

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

type ActualiteMotCle struct {
	Id          int       `json:"id" orm:"column(id)"`
	Mot         string    `json:"mot" orm:"column(mot)"`
	Statut      bool      `json:"statut" orm:"column(statut)"`
	DateUpd     time.Time `json:"date_upd" orm:"column(date_upd);type(timestamp with time zone)"`
	DateAdd     time.Time `json:"date_add" orm:"column(date_add)"`
	CategorieId int64     `json:"categorie_id" orm:"column(categorie_id)"`
}

func (t *ActualiteMotCle) TableName() string {
	return "actualite_article"
}

func init() {
	orm.RegisterModel(&ActualiteMotCle{})
}

var MotCle = graphql.NewObject(graphql.ObjectConfig{
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

// AddActualiteMotCle insert a new ActualiteMotCle into database and returns
// last inserted Id on success.
func AddActualiteMotCle(m *ActualiteMotCle) (Iscreated bool, id int64, err error) {
	o := orm.NewOrm()
	Iscreated, id, err = o.ReadOrCreate(m, "mot")
	return
}

//SelectActualiteMotCle Select all ActualiteMotCle into the database and returns
//[]ActualiteArticel
func SelectAllMotCle() ([]*ActualiteMotCle, error) {
	o := orm.NewOrm()
	var motCle []*ActualiteMotCle
	_, err := o.QueryTable(&ActualiteMotCle{}).All(&motCle)
	return motCle, err
}
