package db

import (
	"fmt"
	"runtime"

	"github.com/201RichK/graphql/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Conn() (*gorm.DB, error) {
	conf := utils.Conf()

	Db, err = gorm.Open("postgres", fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=%s", conf.Server.Host, conf.Database.Username, conf.Database.DbPort, conf.Database.Password, conf.Database.Dbname, conf.Database.Sslmode))
	if err != nil {
		logrus.Error("Gorm Database connection failed to open")
		return nil, err
	}

	return Db, nil
}
