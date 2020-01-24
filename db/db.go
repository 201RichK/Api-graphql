package db

import (
	"runtime"
 	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)


func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:password@localhost/media?sslmode=disable")

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
	orm.Debug = false
}