package db

import (
	"runtime"

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

	Db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=media password=password sslmode=disable")
	if err != nil {
		logrus.Error("Gorm Database connection failed to open")
		panic(err)
	}

	logrus.Info("Database connection established")
}
